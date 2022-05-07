// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Stats Stats provides the overall progress of the scan all process.
//
// swagger:model Stats
type Stats struct {

	// The number of the finished scan processes triggered by the scan all action
	// Example: 90
	Completed int64 `json:"completed"`

	// The metrics data for the each status
	// Example: {"Error":2,"Running":3,"Success":5}
	Metrics map[string]int64 `json:"metrics,omitempty"`

	// A flag indicating job status of scan all.
	Ongoing bool `json:"ongoing"`

	// The total number of scan processes triggered by the scan all action
	// Example: 100
	Total int64 `json:"total"`

	// The trigger of the scan all job.
	// Enum: [Manual Schedule Event]
	Trigger string `json:"trigger,omitempty"`
}

// Validate validates this stats
func (m *Stats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTrigger(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var statsTypeTriggerPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Manual","Schedule","Event"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		statsTypeTriggerPropEnum = append(statsTypeTriggerPropEnum, v)
	}
}

const (

	// StatsTriggerManual captures enum value "Manual"
	StatsTriggerManual string = "Manual"

	// StatsTriggerSchedule captures enum value "Schedule"
	StatsTriggerSchedule string = "Schedule"

	// StatsTriggerEvent captures enum value "Event"
	StatsTriggerEvent string = "Event"
)

// prop value enum
func (m *Stats) validateTriggerEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, statsTypeTriggerPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Stats) validateTrigger(formats strfmt.Registry) error {
	if swag.IsZero(m.Trigger) { // not required
		return nil
	}

	// value enum
	if err := m.validateTriggerEnum("trigger", "body", m.Trigger); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this stats based on context it is used
func (m *Stats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Stats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Stats) UnmarshalBinary(b []byte) error {
	var res Stats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
