package goplugin

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"sync"
	"time"

	"github.com/hashicorp/go-plugin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-common/pkg/services"
	"github.com/smartcontractkit/chainlink-common/pkg/utils"
)

const KeepAliveTickDuration = 5 * time.Second // TODO from config

type grpcPlugin interface {
	plugin.Plugin
	plugin.GRPCPlugin
	ClientConfig() *plugin.ClientConfig
}

// PluginService is a [types.Service] wrapper that maintains an internal [types.Service] created from a [grpcPlugin]
// client instance by launching and re-launching as necessary.
type PluginService[P grpcPlugin, S services.Service] struct {
	services.StateMachine

	pluginName string

	lggr logger.Logger
	cmd  func() *exec.Cmd

	wg     sync.WaitGroup
	stopCh services.StopChan

	grpcPlug P

	client         *plugin.Client
	clientProtocol plugin.ClientProtocol

	newService NewService[S]

	serviceCh      chan struct{} // closed when service is available
	Service        S
	HealthReporter services.HealthReporter // may be the same as Service

	testInterrupt chan func(*PluginService[P, S]) // tests only (via TestHook) to enable access to internals without racing
}

// NewService funcs returns an S and a HeathReporter, which should provide the top level health report for the whole
// plugin, and which may be the same as S.
type NewService[S services.Service] func(context.Context, any) (S, services.HealthReporter, error)

// Init initializes s and should be called from the constructor of the type that embeds it.
//
// newService transforms the type-less plugin in to a Service and HealthReporter. If the plugin is only type-cast to S
// then it should be returned as the HealthReporter too. If additional calls are made after casting to create S, then
// the originally cast value should be returned as the HealthReporter to provide a root level health report.
func (s *PluginService[P, S]) Init(
	pluginName string,
	grpcPlug P,
	newService NewService[S],
	lggr logger.Logger,
	cmd func() *exec.Cmd,
	stopCh chan struct{},
) {
	s.pluginName = pluginName
	s.lggr = lggr
	s.cmd = cmd
	s.stopCh = stopCh
	s.grpcPlug = grpcPlug
	s.newService = newService
	s.serviceCh = make(chan struct{})
}

func (s *PluginService[P, S]) keepAlive() {
	defer s.wg.Done()

	s.lggr.Debugw("Starting keepAlive", "tick", KeepAliveTickDuration)

	check := func() {
		c := s.client
		cp := s.clientProtocol
		if c != nil && !c.Exited() && cp != nil {
			// launched
			err := cp.Ping()
			if err == nil {
				return // healthy
			}
			s.lggr.Errorw("Relaunching unhealthy plugin", "err", err)
		}
		if err := s.tryLaunch(cp); err != nil {
			s.lggr.Errorw("Failed to launch plugin", "err", err)
		}
	}

	check() // no delay

	t := time.NewTicker(KeepAliveTickDuration)
	defer t.Stop()
	for {
		select {
		case <-s.stopCh:
			return
		case <-t.C:
			check()
		case fn := <-s.testInterrupt:
			fn(s)
		}
	}
}

func (s *PluginService[P, S]) tryLaunch(old plugin.ClientProtocol) (err error) {
	if old != nil && s.clientProtocol != old {
		// already replaced by another routine
		return nil
	}
	if cerr := s.closeClient(); cerr != nil {
		s.lggr.Errorw("Error closing old client", "err", cerr)
	}
	s.client, s.clientProtocol, err = s.launch()
	return
}

func (s *PluginService[P, S]) launch() (*plugin.Client, plugin.ClientProtocol, error) {
	ctx, cancelFn := utils.ContextFromChan(s.stopCh)
	defer cancelFn()

	s.lggr.Debug("Launching")

	cc := s.grpcPlug.ClientConfig()
	cc.SkipHostEnv = true
	cc.Cmd = s.cmd()
	client := plugin.NewClient(cc)
	cp, err := client.Client()
	if err != nil {
		client.Kill()
		return nil, nil, fmt.Errorf("failed to create ClientProtocol: %w", err)
	}
	abort := func() {
		if cerr := cp.Close(); cerr != nil {
			s.lggr.Errorw("Error closing ClientProtocol", "err", cerr)
		}
		client.Kill()
	}
	i, err := cp.Dispense(s.pluginName)
	if err != nil {
		abort()
		return nil, nil, fmt.Errorf("failed to Dispense %q plugin: %w", s.pluginName, err)
	}

	select {
	case <-s.serviceCh:
		// s.service already set
	default:
		s.Service, s.HealthReporter, err = s.newService(ctx, i)
		if err != nil {
			abort()
			return nil, nil, fmt.Errorf("failed to create service: %w", err)
		}
		defer close(s.serviceCh)
	}
	return client, cp, nil
}

func (s *PluginService[P, S]) Start(context.Context) error {
	return s.StartOnce("PluginService", func() error {
		s.wg.Add(1)
		go s.keepAlive()
		return nil
	})
}

func (s *PluginService[P, S]) Ready() error {
	select {
	case <-s.serviceCh:
		return s.Service.Ready()
	default:
		return ErrPluginUnavailable
	}
}

func (s *PluginService[P, S]) Name() string { return s.lggr.Name() }

func (s *PluginService[P, S]) HealthReport() map[string]error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	select {
	case <-s.stopCh:
		return map[string]error{s.Name(): fmt.Errorf("service was stoped while waiting: %w", context.Canceled)}
	case <-ctx.Done():
		return map[string]error{s.Name(): ErrPluginUnavailable}
	case <-s.serviceCh:
	}

	hr := map[string]error{s.Name(): s.Healthy()}

	// wait until service is ready, which also triggers the deferred construction to ensure a complete HealthReport
	err := s.Service.Ready()
	for err != nil {
		select {
		case <-s.stopCh:
			return map[string]error{s.Name(): fmt.Errorf("service was stoped while waiting: %w", context.Canceled)}
		case <-ctx.Done():
			hr[s.Service.Name()] = err
			return hr
		case <-time.After(time.Second):
			err = s.Service.Ready()
		}
	}

	services.CopyHealth(hr, s.HealthReporter.HealthReport())
	return hr
}

func (s *PluginService[P, S]) Close() error {
	return s.StopOnce("PluginService", func() (err error) {
		close(s.stopCh)
		s.wg.Wait()

		select {
		case <-s.serviceCh:
			if cerr := s.Service.Close(); !isCanceled(cerr) {
				err = errors.Join(err, cerr)
			}
		default:
		}
		err = errors.Join(err, s.closeClient())
		return
	})
}

func (s *PluginService[P, S]) closeClient() (err error) {
	if s.clientProtocol != nil {
		if cerr := s.clientProtocol.Close(); !isCanceled(cerr) {
			err = cerr
		}
	}
	if s.client != nil {
		s.client.Kill()
	}
	return
}

// WaitCtx waits for the service to start up until `ctx.Done` is triggered
// or it receives the stop signal.
func (s *PluginService[P, S]) WaitCtx(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return context.Cause(ctx)
	case <-s.serviceCh:
		return nil
	case <-s.stopCh:
		return fmt.Errorf("service was stoped while waiting: %w", context.Canceled)
	}
}

// XXXTestHook returns a TestPluginService.
// It must only be called once, and before Start.
func (s *PluginService[P, S]) XXXTestHook() TestPluginService[P, S] {
	s.testInterrupt = make(chan func(*PluginService[P, S]))
	return s.testInterrupt
}

// TestPluginService supports Killing & Resetting a running *pluginService.
type TestPluginService[P grpcPlugin, S services.Service] chan<- func(*PluginService[P, S])

func (ch TestPluginService[P, S]) Kill() {
	done := make(chan struct{})
	ch <- func(s *PluginService[P, S]) {
		defer close(done)
		_ = s.closeClient()
	}
	<-done
}

func (ch TestPluginService[P, S]) Reset() {
	done := make(chan struct{})
	ch <- func(r *PluginService[P, S]) {
		defer close(done)
		_ = r.closeClient()
		r.client = nil
		r.clientProtocol = nil
	}
	<-done
}

func isCanceled(err error) bool {
	return errors.Is(err, context.Canceled) || status.Code(err) == codes.Canceled
}
