// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// GetPolicyReader is a Reader for the GetPolicy structure.
type GetPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPolicyBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPolicyUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPolicyForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPolicyNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPolicyInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetPolicyOK creates a GetPolicyOK with default headers values
func NewGetPolicyOK() *GetPolicyOK {
	return &GetPolicyOK{}
}

/* GetPolicyOK describes a response with status code 200, with default header values.

Get a preheat policy success
*/
type GetPolicyOK struct {
	Payload *model.PreheatPolicy
}

func (o *GetPolicyOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyOK  %+v", 200, o.Payload)
}
func (o *GetPolicyOK) GetPayload() *model.PreheatPolicy {
	return o.Payload
}

func (o *GetPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.PreheatPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPolicyBadRequest creates a GetPolicyBadRequest with default headers values
func NewGetPolicyBadRequest() *GetPolicyBadRequest {
	return &GetPolicyBadRequest{}
}

/* GetPolicyBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetPolicyBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPolicyBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyBadRequest  %+v", 400, o.Payload)
}
func (o *GetPolicyBadRequest) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPolicyBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPolicyUnauthorized creates a GetPolicyUnauthorized with default headers values
func NewGetPolicyUnauthorized() *GetPolicyUnauthorized {
	return &GetPolicyUnauthorized{}
}

/* GetPolicyUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetPolicyUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPolicyUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyUnauthorized  %+v", 401, o.Payload)
}
func (o *GetPolicyUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPolicyUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPolicyForbidden creates a GetPolicyForbidden with default headers values
func NewGetPolicyForbidden() *GetPolicyForbidden {
	return &GetPolicyForbidden{}
}

/* GetPolicyForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetPolicyForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPolicyForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyForbidden  %+v", 403, o.Payload)
}
func (o *GetPolicyForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPolicyForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPolicyNotFound creates a GetPolicyNotFound with default headers values
func NewGetPolicyNotFound() *GetPolicyNotFound {
	return &GetPolicyNotFound{}
}

/* GetPolicyNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetPolicyNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPolicyNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyNotFound  %+v", 404, o.Payload)
}
func (o *GetPolicyNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPolicyNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetPolicyInternalServerError creates a GetPolicyInternalServerError with default headers values
func NewGetPolicyInternalServerError() *GetPolicyInternalServerError {
	return &GetPolicyInternalServerError{}
}

/* GetPolicyInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetPolicyInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetPolicyInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name}/preheat/policies/{preheat_policy_name}][%d] getPolicyInternalServerError  %+v", 500, o.Payload)
}
func (o *GetPolicyInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetPolicyInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
