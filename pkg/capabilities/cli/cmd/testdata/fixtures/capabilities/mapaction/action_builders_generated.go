// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package mapaction

import (
	"github.com/smartcontractkit/chainlink-common/pkg/capabilities"
	"github.com/smartcontractkit/chainlink-common/pkg/workflows/sdk"
)

func (cfg ActionConfig) New(w *sdk.WorkflowSpecFactory, ref string, input ActionInput) ActionOutputsCap {

	def := sdk.StepDefinition{
		ID: "mapaction@1.0.0", Ref: ref,
		Inputs:         input.ToSteps(),
		Config:         map[string]any{},
		CapabilityType: capabilities.CapabilityTypeAction,
	}

	step := sdk.Step[ActionOutputs]{Definition: def}
	return ActionOutputsCapFromStep(w, step)
}

type ActionOutputsCap interface {
	sdk.CapDefinition[ActionOutputs]
	Payload() ActionOutputsPayloadCap
	private()
}

// ActionOutputsCapFromStep should only be called from generated code to assure type safety
func ActionOutputsCapFromStep(w *sdk.WorkflowSpecFactory, step sdk.Step[ActionOutputs]) ActionOutputsCap {
	raw := step.AddTo(w)
	return &actionOutputs{CapDefinition: raw}
}

type actionOutputs struct {
	sdk.CapDefinition[ActionOutputs]
}

func (*actionOutputs) private() {}
func (c *actionOutputs) Payload() ActionOutputsPayloadCap {
	return ActionOutputsPayloadCap(sdk.AccessField[ActionOutputs, ActionOutputsPayload](c.CapDefinition, "payload"))
}

func NewActionOutputsFromFields(
	payload ActionOutputsPayloadCap) ActionOutputsCap {
	return &simpleActionOutputs{
		CapDefinition: sdk.ComponentCapDefinition[ActionOutputs]{
			"payload": payload.Ref(),
		},
		payload: payload,
	}
}

type simpleActionOutputs struct {
	sdk.CapDefinition[ActionOutputs]
	payload ActionOutputsPayloadCap
}

func (c *simpleActionOutputs) Payload() ActionOutputsPayloadCap {
	return c.payload
}

func (c *simpleActionOutputs) private() {}

type ActionOutputsPayloadCap sdk.CapDefinition[ActionOutputsPayload]

type ActionInput struct {
	Payload sdk.CapDefinition[ActionInputsPayload]
}

func (input ActionInput) ToSteps() sdk.StepInputs {
	return sdk.StepInputs{
		Mapping: map[string]any{
			"payload": input.Payload.Ref(),
		},
	}
}
