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

// GetSearchReader is a Reader for the GetSearch structure.
type GetSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetSearchInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetSearchOK creates a GetSearchOK with default headers values
func NewGetSearchOK() *GetSearchOK {
	return &GetSearchOK{}
}

/* GetSearchOK describes a response with status code 200, with default header values.

An array of search results
*/
type GetSearchOK struct {
	Payload []*legacy.Search
}

func (o *GetSearchOK) Error() string {
	return fmt.Sprintf("[GET /search][%d] getSearchOK  %+v", 200, o.Payload)
}
func (o *GetSearchOK) GetPayload() []*legacy.Search {
	return o.Payload
}

func (o *GetSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSearchInternalServerError creates a GetSearchInternalServerError with default headers values
func NewGetSearchInternalServerError() *GetSearchInternalServerError {
	return &GetSearchInternalServerError{}
}

/* GetSearchInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetSearchInternalServerError struct {
}

func (o *GetSearchInternalServerError) Error() string {
	return fmt.Sprintf("[GET /search][%d] getSearchInternalServerError ", 500)
}

func (o *GetSearchInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
