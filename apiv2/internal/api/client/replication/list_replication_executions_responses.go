// Code generated by go-swagger; DO NOT EDIT.

package replication

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zelat/goharbor-client/apiv2/model"
)

// ListReplicationExecutionsReader is a Reader for the ListReplicationExecutions structure.
type ListReplicationExecutionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListReplicationExecutionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListReplicationExecutionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewListReplicationExecutionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListReplicationExecutionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListReplicationExecutionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewListReplicationExecutionsOK creates a ListReplicationExecutionsOK with default headers values
func NewListReplicationExecutionsOK() *ListReplicationExecutionsOK {
	return &ListReplicationExecutionsOK{}
}

/* ListReplicationExecutionsOK describes a response with status code 200, with default header values.

Success
*/
type ListReplicationExecutionsOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of the resources
	 */
	XTotalCount int64

	Payload []*model.ReplicationExecution
}

func (o *ListReplicationExecutionsOK) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] listReplicationExecutionsOK  %+v", 200, o.Payload)
}
func (o *ListReplicationExecutionsOK) GetPayload() []*model.ReplicationExecution {
	return o.Payload
}

func (o *ListReplicationExecutionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListReplicationExecutionsUnauthorized creates a ListReplicationExecutionsUnauthorized with default headers values
func NewListReplicationExecutionsUnauthorized() *ListReplicationExecutionsUnauthorized {
	return &ListReplicationExecutionsUnauthorized{}
}

/* ListReplicationExecutionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type ListReplicationExecutionsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationExecutionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] listReplicationExecutionsUnauthorized  %+v", 401, o.Payload)
}
func (o *ListReplicationExecutionsUnauthorized) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationExecutionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListReplicationExecutionsForbidden creates a ListReplicationExecutionsForbidden with default headers values
func NewListReplicationExecutionsForbidden() *ListReplicationExecutionsForbidden {
	return &ListReplicationExecutionsForbidden{}
}

/* ListReplicationExecutionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type ListReplicationExecutionsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationExecutionsForbidden) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] listReplicationExecutionsForbidden  %+v", 403, o.Payload)
}
func (o *ListReplicationExecutionsForbidden) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationExecutionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewListReplicationExecutionsInternalServerError creates a ListReplicationExecutionsInternalServerError with default headers values
func NewListReplicationExecutionsInternalServerError() *ListReplicationExecutionsInternalServerError {
	return &ListReplicationExecutionsInternalServerError{}
}

/* ListReplicationExecutionsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type ListReplicationExecutionsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *model.Errors
}

func (o *ListReplicationExecutionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /replication/executions][%d] listReplicationExecutionsInternalServerError  %+v", 500, o.Payload)
}
func (o *ListReplicationExecutionsInternalServerError) GetPayload() *model.Errors {
	return o.Payload
}

func (o *ListReplicationExecutionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
