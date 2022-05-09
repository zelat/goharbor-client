package repository

import (
	"net/http"

	"github.com/go-openapi/runtime"
	"github.com/zelat/goharbor-client/apiv2/internal/legacyapi/client/products"
)

const (
	// ErrRegistryIllegalIDFormat describes an illegal request format
	ErrRepositoryIllegalIDFormatMsg = "illegal format of provided ID value"
	// ErrRegistryUnauthorized describes an unauthorized request
	ErrRepositoryUnauthorizedMsg = "unauthorized"
	// ErrRegistryInternalErrors describes server-side internal errors
	ErrRepositoryInternalErrorsMsg = "unexpected internal errors"
	// ErrRegistryNoPermission describes a request error without permission
	ErrRepositoryNoPermissionMsg = "user does not have permission to the registry"
	// ErrRegistryIDNotExists describes an error
	// when no proper registry ID is found
	ErrRepositoryIDNotExistsMsg = "registry ID does not exist"
	// ErrRegistryNameAlreadyExists describes a duplicate registry name error
	ErrRepositoryNameAlreadyExistsMsg = "registry name already exists"
	// ErrRegistryMismatch describes a failed lookup
	// of a registry with name/id pair
	ErrRepositoryMismatchMsg = "id/name pair not found on server side"
	// ErrRegistryNotFound describes an error
	// when a specific registry is not found
	ErrRepositoryNotFoundMsg = "registry not found on server side"
	// ErrRegistryNotProvidedMsg describes an error
	// when no registry was provided by the user
	ErrRepositoryNotProvidedMsg = "no registry provided"
)

// ErrRegistryIllegalIDFormat describes an illegal request format.
type ErrRepositoryIllegalIDFormat struct{}

// Error returns the error message.
func (e *ErrRepositoryIllegalIDFormat) Error() string {
	return ErrRepositoryIllegalIDFormatMsg
}

// ErrRegistryUnauthorized describes an unauthorized request.
type ErrRepositoryUnauthorized struct{}

// Error returns the error message.
func (e *ErrRepositoryUnauthorized) Error() string {
	return ErrRepositoryUnauthorizedMsg
}

// ErrRegistryInternalErrors describes server-side internal errors.
type ErrRepositoryInternalErrors struct{}

// Error returns the error message.
func (e *ErrRepositoryInternalErrors) Error() string {
	return ErrRepositoryInternalErrorsMsg
}

// ErrRegistryNoPermission describes a request error without permission.
type ErrRepositoryNoPermission struct{}

// Error returns the error message.
func (e *ErrRepositoryNoPermission) Error() string {
	return ErrRepositoryNoPermissionMsg
}

// ErrRegistryIDNotExists describes an error
// when no proper registry ID is found.
type ErrRepositoryIDNotExists struct{}

// Error returns the error message.
func (e *ErrRepositoryIDNotExists) Error() string {
	return ErrRepositoryIDNotExistsMsg
}

// ErrRegistryNameAlreadyExists describes a duplicate registry name error.
type ErrRepositoryNameAlreadyExists struct{}

// Error returns the error message.
func (e *ErrRepositoryNameAlreadyExists) Error() string {
	return ErrRepositoryNameAlreadyExistsMsg
}

// ErrRegistryMismatch describes a failed lookup
// of a registry with name/id pair.
type ErrRepositoryMismatch struct{}

// Error returns the error message.
func (e *ErrRepositoryMismatch) Error() string {
	return ErrRepositoryMismatchMsg
}

// ErrRegistryNotFound describes an error
// when a specific registry is not found.
type ErrRepositoryNotFound struct{}

// Error returns the error message.
func (e *ErrRepositoryNotFound) Error() string {
	return ErrRepositoryNotFoundMsg
}

// ErrRegistryNotProvided describes an error
// when no registry was provided.
type ErrRepositoryNotProvided struct{}

// Error returns the error message.
func (e *ErrRepositoryNotProvided) Error() string {
	return ErrRepositoryNotProvidedMsg
}

// handleRegistryErrors takes a swagger generated error as input,
// which usually does not contain any form of error message,
// and outputs a new error with proper message.
func handleSwaggerRepositoryErrors(in error) error {
	t, ok := in.(*runtime.APIError)
	if ok {
		switch t.Code {
		case http.StatusBadRequest:
			return &ErrRepositoryIllegalIDFormat{}
		case http.StatusUnauthorized:
			return &ErrRepositoryUnauthorized{}
		case http.StatusForbidden:
			return &ErrRepositoryNoPermission{}
		case http.StatusInternalServerError:
			return &ErrRepositoryInternalErrors{}
		}
	}

	switch in.(type) {
	case *products.DeleteRegistriesIDNotFound:
		return &ErrRepositoryIDNotExists{}
	case *products.PutRegistriesIDNotFound:
		return &ErrRepositoryIDNotExists{}
	case *products.PostRegistriesConflict:
		return &ErrRepositoryNameAlreadyExists{}
	default:
		return in
	}
}
