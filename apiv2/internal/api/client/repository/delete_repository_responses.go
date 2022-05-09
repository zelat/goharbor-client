// Code generated by go-swagger; DO NOT EDIT.

package repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// DeleteRepositoryReader is a Reader for the DeleteRepository structure.
type DeleteRepositoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRepositoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRepositoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRepositoryBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRepositoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRepositoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRepositoryNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRepositoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRepositoryOK creates a DeleteRepositoryOK with default headers values
func NewDeleteRepositoryOK() *DeleteRepositoryOK {
	return &DeleteRepositoryOK{}
}

/* DeleteRepositoryOK describes a response with status code 200, with default header values.

Success
*/
type DeleteRepositoryOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteRepositoryOK) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryOK ", 200)
}

func (o *DeleteRepositoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteRepositoryBadRequest creates a DeleteRepositoryBadRequest with default headers values
func NewDeleteRepositoryBadRequest() *DeleteRepositoryBadRequest {
	return &DeleteRepositoryBadRequest{}
}

/* DeleteRepositoryBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteRepositoryBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRepositoryBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteRepositoryBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRepositoryBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepositoryUnauthorized creates a DeleteRepositoryUnauthorized with default headers values
func NewDeleteRepositoryUnauthorized() *DeleteRepositoryUnauthorized {
	return &DeleteRepositoryUnauthorized{}
}

/* DeleteRepositoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRepositoryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRepositoryUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteRepositoryUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRepositoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepositoryForbidden creates a DeleteRepositoryForbidden with default headers values
func NewDeleteRepositoryForbidden() *DeleteRepositoryForbidden {
	return &DeleteRepositoryForbidden{}
}

/* DeleteRepositoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRepositoryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRepositoryForbidden) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryForbidden  %+v", 403, o.Payload)
}
func (o *DeleteRepositoryForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRepositoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepositoryNotFound creates a DeleteRepositoryNotFound with default headers values
func NewDeleteRepositoryNotFound() *DeleteRepositoryNotFound {
	return &DeleteRepositoryNotFound{}
}

/* DeleteRepositoryNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteRepositoryNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRepositoryNotFound) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryNotFound  %+v", 404, o.Payload)
}
func (o *DeleteRepositoryNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRepositoryNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteRepositoryInternalServerError creates a DeleteRepositoryInternalServerError with default headers values
func NewDeleteRepositoryInternalServerError() *DeleteRepositoryInternalServerError {
	return &DeleteRepositoryInternalServerError{}
}

/* DeleteRepositoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteRepositoryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRepositoryInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /projects/{project_name}/repositories/{repository_name}][%d] deleteRepositoryInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteRepositoryInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRepositoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(model.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
