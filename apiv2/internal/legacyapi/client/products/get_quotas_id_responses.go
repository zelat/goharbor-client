// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/v4/apiv2/model/legacy"
)

// GetQuotasIDReader is a Reader for the GetQuotasID structure.
type GetQuotasIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetQuotasIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetQuotasIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetQuotasIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetQuotasIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetQuotasIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetQuotasIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetQuotasIDOK creates a GetQuotasIDOK with default headers values
func NewGetQuotasIDOK() *GetQuotasIDOK {
	return &GetQuotasIDOK{}
}

/* GetQuotasIDOK describes a response with status code 200, with default header values.

Successfully retrieved the quota.
*/
type GetQuotasIDOK struct {
	Payload *legacy.Quota
}

func (o *GetQuotasIDOK) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotasIdOK  %+v", 200, o.Payload)
}
func (o *GetQuotasIDOK) GetPayload() *legacy.Quota {
	return o.Payload
}

func (o *GetQuotasIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(legacy.Quota)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetQuotasIDUnauthorized creates a GetQuotasIDUnauthorized with default headers values
func NewGetQuotasIDUnauthorized() *GetQuotasIDUnauthorized {
	return &GetQuotasIDUnauthorized{}
}

/* GetQuotasIDUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetQuotasIDUnauthorized struct {
}

func (o *GetQuotasIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotasIdUnauthorized ", 401)
}

func (o *GetQuotasIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQuotasIDForbidden creates a GetQuotasIDForbidden with default headers values
func NewGetQuotasIDForbidden() *GetQuotasIDForbidden {
	return &GetQuotasIDForbidden{}
}

/* GetQuotasIDForbidden describes a response with status code 403, with default header values.

User does not have permission to call this API
*/
type GetQuotasIDForbidden struct {
}

func (o *GetQuotasIDForbidden) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotasIdForbidden ", 403)
}

func (o *GetQuotasIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQuotasIDNotFound creates a GetQuotasIDNotFound with default headers values
func NewGetQuotasIDNotFound() *GetQuotasIDNotFound {
	return &GetQuotasIDNotFound{}
}

/* GetQuotasIDNotFound describes a response with status code 404, with default header values.

Quota does not exist.
*/
type GetQuotasIDNotFound struct {
}

func (o *GetQuotasIDNotFound) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotasIdNotFound ", 404)
}

func (o *GetQuotasIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetQuotasIDInternalServerError creates a GetQuotasIDInternalServerError with default headers values
func NewGetQuotasIDInternalServerError() *GetQuotasIDInternalServerError {
	return &GetQuotasIDInternalServerError{}
}

/* GetQuotasIDInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetQuotasIDInternalServerError struct {
}

func (o *GetQuotasIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /quotas/{id}][%d] getQuotasIdInternalServerError ", 500)
}

func (o *GetQuotasIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
