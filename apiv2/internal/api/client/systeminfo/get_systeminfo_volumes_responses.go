// Code generated by go-swagger; DO NOT EDIT.

package systeminfo

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// GetSysteminfoVolumesReader is a Reader for the GetSysteminfoVolumes structure.
type GetSysteminfoVolumesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSysteminfoVolumesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSysteminfoVolumesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSysteminfoVolumesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSysteminfoVolumesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSysteminfoVolumesNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSysteminfoVolumesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSysteminfoVolumesOK creates a GetSysteminfoVolumesOK with default headers values
func NewGetSysteminfoVolumesOK() *GetSysteminfoVolumesOK {
	return &GetSysteminfoVolumesOK{}
}

/* GetSysteminfoVolumesOK describes a response with status code 200, with default header values.

Get system volumes successfully.
*/
type GetSysteminfoVolumesOK struct {
	Payload *model.SystemInfo
}

func (o *GetSysteminfoVolumesOK) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesOK  %+v", 200, o.Payload)
}
func (o *GetSysteminfoVolumesOK) GetPayload() *model.SystemInfo {
	return o.Payload
}

func (o *GetSysteminfoVolumesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.SystemInfo)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSysteminfoVolumesUnauthorized creates a GetSysteminfoVolumesUnauthorized with default headers values
func NewGetSysteminfoVolumesUnauthorized() *GetSysteminfoVolumesUnauthorized {
	return &GetSysteminfoVolumesUnauthorized{}
}

/* GetSysteminfoVolumesUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetSysteminfoVolumesUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSysteminfoVolumesUnauthorized) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesUnauthorized  %+v", 401, o.Payload)
}
func (o *GetSysteminfoVolumesUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSysteminfoVolumesUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSysteminfoVolumesForbidden creates a GetSysteminfoVolumesForbidden with default headers values
func NewGetSysteminfoVolumesForbidden() *GetSysteminfoVolumesForbidden {
	return &GetSysteminfoVolumesForbidden{}
}

/* GetSysteminfoVolumesForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetSysteminfoVolumesForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSysteminfoVolumesForbidden) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesForbidden  %+v", 403, o.Payload)
}
func (o *GetSysteminfoVolumesForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSysteminfoVolumesForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSysteminfoVolumesNotFound creates a GetSysteminfoVolumesNotFound with default headers values
func NewGetSysteminfoVolumesNotFound() *GetSysteminfoVolumesNotFound {
	return &GetSysteminfoVolumesNotFound{}
}

/* GetSysteminfoVolumesNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetSysteminfoVolumesNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSysteminfoVolumesNotFound) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesNotFound  %+v", 404, o.Payload)
}
func (o *GetSysteminfoVolumesNotFound) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSysteminfoVolumesNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetSysteminfoVolumesInternalServerError creates a GetSysteminfoVolumesInternalServerError with default headers values
func NewGetSysteminfoVolumesInternalServerError() *GetSysteminfoVolumesInternalServerError {
	return &GetSysteminfoVolumesInternalServerError{}
}

/* GetSysteminfoVolumesInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetSysteminfoVolumesInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *GetSysteminfoVolumesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /systeminfo/volumes][%d] getSysteminfoVolumesInternalServerError  %+v", 500, o.Payload)
}
func (o *GetSysteminfoVolumesInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *GetSysteminfoVolumesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
