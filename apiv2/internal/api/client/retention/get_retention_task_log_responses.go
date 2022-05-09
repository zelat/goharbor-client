// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// GetRetentionTaskLogReader is a Reader for the GetRetentionTaskLog structure.
type GetRetentionTaskLogReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRetentionTaskLogReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRetentionTaskLogOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRetentionTaskLogUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRetentionTaskLogForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRetentionTaskLogInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetRetentionTaskLogOK creates a GetRetentionTaskLogOK with default headers values
func NewGetRetentionTaskLogOK() *GetRetentionTaskLogOK {
	return &GetRetentionTaskLogOK{}
}

/* GetRetentionTaskLogOK describes a response with status code 200, with default header values.

Get Retention job task log successfully.
*/
type GetRetentionTaskLogOK struct {
	Payload string
}

func (o *GetRetentionTaskLogOK) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks/{tid}][%d] getRetentionTaskLogOK  %+v", 200, o.Payload)
}
func (o *GetRetentionTaskLogOK) GetPayload() string {
	return o.Payload
}

func (o *GetRetentionTaskLogOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRetentionTaskLogUnauthorized creates a GetRetentionTaskLogUnauthorized with default headers values
func NewGetRetentionTaskLogUnauthorized() *GetRetentionTaskLogUnauthorized {
	return &GetRetentionTaskLogUnauthorized{}
}

/* GetRetentionTaskLogUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRetentionTaskLogUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRetentionTaskLogUnauthorized) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks/{tid}][%d] getRetentionTaskLogUnauthorized  %+v", 401, o.Payload)
}
func (o *GetRetentionTaskLogUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRetentionTaskLogUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRetentionTaskLogForbidden creates a GetRetentionTaskLogForbidden with default headers values
func NewGetRetentionTaskLogForbidden() *GetRetentionTaskLogForbidden {
	return &GetRetentionTaskLogForbidden{}
}

/* GetRetentionTaskLogForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRetentionTaskLogForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRetentionTaskLogForbidden) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks/{tid}][%d] getRetentionTaskLogForbidden  %+v", 403, o.Payload)
}
func (o *GetRetentionTaskLogForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRetentionTaskLogForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetRetentionTaskLogInternalServerError creates a GetRetentionTaskLogInternalServerError with default headers values
func NewGetRetentionTaskLogInternalServerError() *GetRetentionTaskLogInternalServerError {
	return &GetRetentionTaskLogInternalServerError{}
}

/* GetRetentionTaskLogInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetRetentionTaskLogInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetRetentionTaskLogInternalServerError) Error() string {
	return fmt.Sprintf("[GET /retentions/{id}/executions/{eid}/tasks/{tid}][%d] getRetentionTaskLogInternalServerError  %+v", 500, o.Payload)
}
func (o *GetRetentionTaskLogInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetRetentionTaskLogInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
