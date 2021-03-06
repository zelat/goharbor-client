// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/zelat/goharbor-client/apiv2/model/legacy"
)

// GetProjectsProjectIDWebhookJobsReader is a Reader for the GetProjectsProjectIDWebhookJobs structure.
type GetProjectsProjectIDWebhookJobsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetProjectsProjectIDWebhookJobsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetProjectsProjectIDWebhookJobsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetProjectsProjectIDWebhookJobsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetProjectsProjectIDWebhookJobsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetProjectsProjectIDWebhookJobsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetProjectsProjectIDWebhookJobsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetProjectsProjectIDWebhookJobsOK creates a GetProjectsProjectIDWebhookJobsOK with default headers values
func NewGetProjectsProjectIDWebhookJobsOK() *GetProjectsProjectIDWebhookJobsOK {
	return &GetProjectsProjectIDWebhookJobsOK{}
}

/* GetProjectsProjectIDWebhookJobsOK describes a response with status code 200, with default header values.

List project webhook jobs successfully.
*/
type GetProjectsProjectIDWebhookJobsOK struct {

	/* Link to previous page and next page
	 */
	Link string

	/* The total count of available items
	 */
	XTotalCount int64

	Payload []*legacy.WebhookJob
}

func (o *GetProjectsProjectIDWebhookJobsOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/jobs][%d] getProjectsProjectIdWebhookJobsOK  %+v", 200, o.Payload)
}
func (o *GetProjectsProjectIDWebhookJobsOK) GetPayload() []*legacy.WebhookJob {
	return o.Payload
}

func (o *GetProjectsProjectIDWebhookJobsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetProjectsProjectIDWebhookJobsBadRequest creates a GetProjectsProjectIDWebhookJobsBadRequest with default headers values
func NewGetProjectsProjectIDWebhookJobsBadRequest() *GetProjectsProjectIDWebhookJobsBadRequest {
	return &GetProjectsProjectIDWebhookJobsBadRequest{}
}

/* GetProjectsProjectIDWebhookJobsBadRequest describes a response with status code 400, with default header values.

Illegal format of provided ID value.
*/
type GetProjectsProjectIDWebhookJobsBadRequest struct {
}

func (o *GetProjectsProjectIDWebhookJobsBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/jobs][%d] getProjectsProjectIdWebhookJobsBadRequest ", 400)
}

func (o *GetProjectsProjectIDWebhookJobsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookJobsUnauthorized creates a GetProjectsProjectIDWebhookJobsUnauthorized with default headers values
func NewGetProjectsProjectIDWebhookJobsUnauthorized() *GetProjectsProjectIDWebhookJobsUnauthorized {
	return &GetProjectsProjectIDWebhookJobsUnauthorized{}
}

/* GetProjectsProjectIDWebhookJobsUnauthorized describes a response with status code 401, with default header values.

User need to log in first.
*/
type GetProjectsProjectIDWebhookJobsUnauthorized struct {
}

func (o *GetProjectsProjectIDWebhookJobsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/jobs][%d] getProjectsProjectIdWebhookJobsUnauthorized ", 401)
}

func (o *GetProjectsProjectIDWebhookJobsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookJobsForbidden creates a GetProjectsProjectIDWebhookJobsForbidden with default headers values
func NewGetProjectsProjectIDWebhookJobsForbidden() *GetProjectsProjectIDWebhookJobsForbidden {
	return &GetProjectsProjectIDWebhookJobsForbidden{}
}

/* GetProjectsProjectIDWebhookJobsForbidden describes a response with status code 403, with default header values.

User have no permission to list webhook jobs of the project.
*/
type GetProjectsProjectIDWebhookJobsForbidden struct {
}

func (o *GetProjectsProjectIDWebhookJobsForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/jobs][%d] getProjectsProjectIdWebhookJobsForbidden ", 403)
}

func (o *GetProjectsProjectIDWebhookJobsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetProjectsProjectIDWebhookJobsInternalServerError creates a GetProjectsProjectIDWebhookJobsInternalServerError with default headers values
func NewGetProjectsProjectIDWebhookJobsInternalServerError() *GetProjectsProjectIDWebhookJobsInternalServerError {
	return &GetProjectsProjectIDWebhookJobsInternalServerError{}
}

/* GetProjectsProjectIDWebhookJobsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetProjectsProjectIDWebhookJobsInternalServerError struct {
}

func (o *GetProjectsProjectIDWebhookJobsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_id}/webhook/jobs][%d] getProjectsProjectIdWebhookJobsInternalServerError ", 500)
}

func (o *GetProjectsProjectIDWebhookJobsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
