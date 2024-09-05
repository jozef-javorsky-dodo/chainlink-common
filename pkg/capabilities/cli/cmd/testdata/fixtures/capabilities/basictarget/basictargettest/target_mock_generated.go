// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package basictargettest

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli/cmd/testdata/fixtures/capabilities/basictarget"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/testutils"
)

// Target registers a new capability mock with the runner
func Target(runner *testutils.Runner, fn func(input basictarget.TargetInputs) error) *testutils.TargetMock[basictarget.TargetInputs] {
	mock := testutils.MockTarget[basictarget.TargetInputs]("basic-test-target@1.0.0", fn)
	runner.MockCapability("basic-test-target@1.0.0", nil, mock)
	return mock
}
