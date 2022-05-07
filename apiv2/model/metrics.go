// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Metrics metrics
//
// swagger:model Metrics
type Metrics struct {

	// The count of error task
	ErrorTaskCount int64 `json:"error_task_count,omitempty"`

	// The count of pending task
	PendingTaskCount int64 `json:"pending_task_count,omitempty"`

	// The count of running task
	RunningTaskCount int64 `json:"running_task_count,omitempty"`

	// The count of scheduled task
	ScheduledTaskCount int64 `json:"scheduled_task_count,omitempty"`

	// The count of stopped task
	StoppedTaskCount int64 `json:"stopped_task_count,omitempty"`

	// The count of success task
	SuccessTaskCount int64 `json:"success_task_count,omitempty"`

	// The count of task
	TaskCount int64 `json:"task_count,omitempty"`
}

// Validate validates this metrics
func (m *Metrics) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this metrics based on context it is used
func (m *Metrics) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Metrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Metrics) UnmarshalBinary(b []byte) error {
	var res Metrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}