// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// DeleteRobotReader is a Reader for the DeleteRobot structure.
type DeleteRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDeleteRobotOK creates a DeleteRobotOK with default headers values
func NewDeleteRobotOK() *DeleteRobotOK {
	return &DeleteRobotOK{}
}

/* DeleteRobotOK describes a response with status code 200, with default header values.

Success
*/
type DeleteRobotOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *DeleteRobotOK) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotOK ", 200)
}

func (o *DeleteRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewDeleteRobotBadRequest creates a DeleteRobotBadRequest with default headers values
func NewDeleteRobotBadRequest() *DeleteRobotBadRequest {
	return &DeleteRobotBadRequest{}
}

/* DeleteRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type DeleteRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotBadRequest  %+v", 400, o.Payload)
}
func (o *DeleteRobotBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotUnauthorized creates a DeleteRobotUnauthorized with default headers values
func NewDeleteRobotUnauthorized() *DeleteRobotUnauthorized {
	return &DeleteRobotUnauthorized{}
}

/* DeleteRobotUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteRobotUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotUnauthorized  %+v", 401, o.Payload)
}
func (o *DeleteRobotUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotForbidden creates a DeleteRobotForbidden with default headers values
func NewDeleteRobotForbidden() *DeleteRobotForbidden {
	return &DeleteRobotForbidden{}
}

/* DeleteRobotForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteRobotForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotForbidden) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotForbidden  %+v", 403, o.Payload)
}
func (o *DeleteRobotForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotNotFound creates a DeleteRobotNotFound with default headers values
func NewDeleteRobotNotFound() *DeleteRobotNotFound {
	return &DeleteRobotNotFound{}
}

/* DeleteRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type DeleteRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotNotFound) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotNotFound  %+v", 404, o.Payload)
}
func (o *DeleteRobotNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewDeleteRobotInternalServerError creates a DeleteRobotInternalServerError with default headers values
func NewDeleteRobotInternalServerError() *DeleteRobotInternalServerError {
	return &DeleteRobotInternalServerError{}
}

/* DeleteRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type DeleteRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *DeleteRobotInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /robots/{robot_id}][%d] deleteRobotInternalServerError  %+v", 500, o.Payload)
}
func (o *DeleteRobotInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *DeleteRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
