// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model/legacy"
)

// GetUsersUserIDReader is a Reader for the GetUsersUserID structure.
type GetUsersUserIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersUserIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersUserIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUsersUserIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUsersUserIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUsersUserIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUsersUserIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUsersUserIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetUsersUserIDOK creates a GetUsersUserIDOK with default headers values
func NewGetUsersUserIDOK() *GetUsersUserIDOK {
	return &GetUsersUserIDOK{}
}

/* GetUsersUserIDOK describes a response with status code 200, with default header values.

Get user's profile successfully.
*/
type GetUsersUserIDOK struct {
	Payload *legacy.User
}

func (o *GetUsersUserIDOK) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdOK  %+v", 200, o.Payload)
}
func (o *GetUsersUserIDOK) GetPayload() *legacy.User {
	return o.Payload
}

func (o *GetUsersUserIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUserIDBadRequest creates a GetUsersUserIDBadRequest with default headers values
func NewGetUsersUserIDBadRequest() *GetUsersUserIDBadRequest {
	return &GetUsersUserIDBadRequest{}
}

/* GetUsersUserIDBadRequest describes a response with status code 400, with default header values.

Invalid user ID.
*/
type GetUsersUserIDBadRequest struct {
}

func (o *GetUsersUserIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdBadRequest ", 400)
}

func (o *GetUsersUserIDBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDUnauthorized creates a GetUsersUserIDUnauthorized with default headers values
func NewGetUsersUserIDUnauthorized() *GetUsersUserIDUnauthorized {
	return &GetUsersUserIDUnauthorized{}
}

/* GetUsersUserIDUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetUsersUserIDUnauthorized struct {
}

func (o *GetUsersUserIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdUnauthorized ", 401)
}

func (o *GetUsersUserIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDForbidden creates a GetUsersUserIDForbidden with default headers values
func NewGetUsersUserIDForbidden() *GetUsersUserIDForbidden {
	return &GetUsersUserIDForbidden{}
}

/* GetUsersUserIDForbidden describes a response with status code 403, with default header values.

User does not have permission of admin role.
*/
type GetUsersUserIDForbidden struct {
}

func (o *GetUsersUserIDForbidden) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdForbidden ", 403)
}

func (o *GetUsersUserIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDNotFound creates a GetUsersUserIDNotFound with default headers values
func NewGetUsersUserIDNotFound() *GetUsersUserIDNotFound {
	return &GetUsersUserIDNotFound{}
}

/* GetUsersUserIDNotFound describes a response with status code 404, with default header values.

User ID does not exist.
*/
type GetUsersUserIDNotFound struct {
}

func (o *GetUsersUserIDNotFound) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdNotFound ", 404)
}

func (o *GetUsersUserIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetUsersUserIDInternalServerError creates a GetUsersUserIDInternalServerError with default headers values
func NewGetUsersUserIDInternalServerError() *GetUsersUserIDInternalServerError {
	return &GetUsersUserIDInternalServerError{}
}

/* GetUsersUserIDInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetUsersUserIDInternalServerError struct {
}

func (o *GetUsersUserIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /users/{user_id}][%d] getUsersUserIdInternalServerError ", 500)
}

func (o *GetUsersUserIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
