// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SearchRepository search repository
//
// swagger:model SearchRepository
type SearchRepository struct {

	// The ID of the project that the repository belongs to
	ProjectID int64 `json:"project_id,omitempty"`

	// The name of the project that the repository belongs to
	ProjectName string `json:"project_name,omitempty"`

	// The flag to indicate the publicity of the project that the repository belongs to (1 is public, 0 is not)
	ProjectPublic bool `json:"project_public,omitempty"`

	// The count how many times the repository is pulled
	PullCount int64 `json:"pull_count,omitempty"`

	// The name of the repository
	RepositoryName string `json:"repository_name,omitempty"`

	// The count of tags in the repository
	TagsCount int64 `json:"tags_count,omitempty"`
}

// Validate validates this search repository
func (m *SearchRepository) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this search repository based on context it is used
func (m *SearchRepository) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SearchRepository) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SearchRepository) UnmarshalBinary(b []byte) error {
	var res SearchRepository
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
