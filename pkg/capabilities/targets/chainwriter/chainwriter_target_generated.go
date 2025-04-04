// Code generated by github.com/smartcontractkit/chainlink-common/pkg/capabilities/cli, DO NOT EDIT.

package chainwriter

import (
	"encoding/json"
	"fmt"

	"reflect"
	"regexp"

	ocr3cap "github.com/smartcontractkit/chainlink-common/pkg/capabilities/consensus/ocr3/ocr3cap"
)

// Writes to a blockchain
type Target struct {
	// Config corresponds to the JSON schema field "config".
	Config TargetConfig `json:"config" yaml:"config" mapstructure:"config"`

	// Inputs corresponds to the JSON schema field "inputs".
	Inputs TargetInputs `json:"inputs" yaml:"inputs" mapstructure:"inputs"`
}

type TargetConfig struct {
	// The address to write to.
	Address string `json:"address" yaml:"address" mapstructure:"address"`

	// The step timeout which must be a number expressed in seconds
	CreStepTimeout int64 `json:"cre_step_timeout" yaml:"cre_step_timeout" mapstructure:"cre_step_timeout"`

	// The delta stage which must be a number followed by a time symbol (s for
	// seconds, m for minutes, h for hours, d for days).
	DeltaStage string `json:"deltaStage" yaml:"deltaStage" mapstructure:"deltaStage"`

	// The schedule which must be the string 'oneAtATime'.
	Schedule TargetConfigSchedule `json:"schedule" yaml:"schedule" mapstructure:"schedule"`
}

type TargetConfigSchedule string

const TargetConfigScheduleAllAtOnce TargetConfigSchedule = "allAtOnce"
const TargetConfigScheduleOneAtATime TargetConfigSchedule = "oneAtATime"

var enumValues_TargetConfigSchedule = []interface{}{
	"oneAtATime",
	"allAtOnce",
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetConfigSchedule) UnmarshalJSON(b []byte) error {
	var v string
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	var ok bool
	for _, expected := range enumValues_TargetConfigSchedule {
		if reflect.DeepEqual(v, expected) {
			ok = true
			break
		}
	}
	if !ok {
		return fmt.Errorf("invalid value (expected one of %#v): %#v", enumValues_TargetConfigSchedule, v)
	}
	*j = TargetConfigSchedule(v)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetConfig) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["address"]; raw != nil && !ok {
		return fmt.Errorf("field address in TargetConfig: required")
	}
	if _, ok := raw["cre_step_timeout"]; raw != nil && !ok {
		return fmt.Errorf("field cre_step_timeout in TargetConfig: required")
	}
	if _, ok := raw["deltaStage"]; raw != nil && !ok {
		return fmt.Errorf("field deltaStage in TargetConfig: required")
	}
	if _, ok := raw["schedule"]; raw != nil && !ok {
		return fmt.Errorf("field schedule in TargetConfig: required")
	}
	type Plain TargetConfig
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	if matched, _ := regexp.MatchString("^[0-9]+[smhd]$", string(plain.DeltaStage)); !matched {
		return fmt.Errorf("field %s pattern match: must match %s", "^[0-9]+[smhd]$", "DeltaStage")
	}
	*j = TargetConfig(plain)
	return nil
}

type TargetInputs struct {
	// SignedReport corresponds to the JSON schema field "signed_report".
	SignedReport ocr3cap.SignedReport `json:"signed_report" yaml:"signed_report" mapstructure:"signed_report"`
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *TargetInputs) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["signed_report"]; raw != nil && !ok {
		return fmt.Errorf("field signed_report in TargetInputs: required")
	}
	type Plain TargetInputs
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = TargetInputs(plain)
	return nil
}

// UnmarshalJSON implements json.Unmarshaler.
func (j *Target) UnmarshalJSON(b []byte) error {
	var raw map[string]interface{}
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}
	if _, ok := raw["config"]; raw != nil && !ok {
		return fmt.Errorf("field config in Target: required")
	}
	if _, ok := raw["inputs"]; raw != nil && !ok {
		return fmt.Errorf("field inputs in Target: required")
	}
	type Plain Target
	var plain Plain
	if err := json.Unmarshal(b, &plain); err != nil {
		return err
	}
	*j = Target(plain)
	return nil
}
