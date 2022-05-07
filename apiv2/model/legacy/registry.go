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

// Registry registry
//
// swagger:model Registry
type Registry struct {

	// The create time of the policy.
	CreationTime string `json:"creation_time,omitempty"`

	// credential
	Credential *RegistryCredential `json:"credential,omitempty"`

	// Description of the registry.
	Description string `json:"description,omitempty"`

	// The registry ID.
	ID int64 `json:"id,omitempty"`

	// Whether or not the certificate will be verified when Harbor tries to access the server.
	Insecure bool `json:"insecure,omitempty"`

	// The registry name.
	Name string `json:"name,omitempty"`

	// Health status of the registry.
	Status string `json:"status,omitempty"`

	// Type of the registry, e.g. 'harbor'.
	Type string `json:"type,omitempty"`

	// The update time of the policy.
	UpdateTime string `json:"update_time,omitempty"`

	// The registry URL string.
	URL string `json:"url,omitempty"`
}

// Validate validates this registry
func (m *Registry) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Registry) validateCredential(formats strfmt.Registry) error {
	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this registry based on the context it is used
func (m *Registry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCredential(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Registry) contextValidateCredential(ctx context.Context, formats strfmt.Registry) error {

	if m.Credential != nil {
		if err := m.Credential.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Registry) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Registry) UnmarshalBinary(b []byte) error {
	var res Registry
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
