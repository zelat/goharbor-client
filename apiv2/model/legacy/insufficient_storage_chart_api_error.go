// Code generated by go-swagger; DO NOT EDIT.

package legacy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// InsufficientStorageChartAPIError Insufficient storage
//
// swagger:model InsufficientStorageChartAPIError
type InsufficientStorageChartAPIError struct {
	ChartAPIError
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *InsufficientStorageChartAPIError) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 ChartAPIError
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.ChartAPIError = aO0

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m InsufficientStorageChartAPIError) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 1)

	aO0, err := swag.WriteJSON(m.ChartAPIError)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this insufficient storage chart API error
func (m *InsufficientStorageChartAPIError) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChartAPIError
	if err := m.ChartAPIError.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this insufficient storage chart API error based on the context it is used
func (m *InsufficientStorageChartAPIError) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with ChartAPIError
	if err := m.ChartAPIError.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *InsufficientStorageChartAPIError) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InsufficientStorageChartAPIError) UnmarshalBinary(b []byte) error {
	var res InsufficientStorageChartAPIError
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}