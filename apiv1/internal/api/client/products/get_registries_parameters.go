// Code generated by go-swagger; DO NOT EDIT.

package products

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewGetRegistriesParams creates a new GetRegistriesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegistriesParams() *GetRegistriesParams {
	return &GetRegistriesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegistriesParamsWithTimeout creates a new GetRegistriesParams object
// with the ability to set a timeout on a request.
func NewGetRegistriesParamsWithTimeout(timeout time.Duration) *GetRegistriesParams {
	return &GetRegistriesParams{
		timeout: timeout,
	}
}

// NewGetRegistriesParamsWithContext creates a new GetRegistriesParams object
// with the ability to set a context for a request.
func NewGetRegistriesParamsWithContext(ctx context.Context) *GetRegistriesParams {
	return &GetRegistriesParams{
		Context: ctx,
	}
}

// NewGetRegistriesParamsWithHTTPClient creates a new GetRegistriesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegistriesParamsWithHTTPClient(client *http.Client) *GetRegistriesParams {
	return &GetRegistriesParams{
		HTTPClient: client,
	}
}

/* GetRegistriesParams contains all the parameters to send to the API endpoint
   for the get registries operation.

   Typically these are written to a http.Request.
*/
type GetRegistriesParams struct {

	/* Name.

	   Registry's name.
	*/
	Name *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistriesParams) WithDefaults() *GetRegistriesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get registries params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegistriesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get registries params
func (o *GetRegistriesParams) WithTimeout(timeout time.Duration) *GetRegistriesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get registries params
func (o *GetRegistriesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get registries params
func (o *GetRegistriesParams) WithContext(ctx context.Context) *GetRegistriesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get registries params
func (o *GetRegistriesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get registries params
func (o *GetRegistriesParams) WithHTTPClient(client *http.Client) *GetRegistriesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get registries params
func (o *GetRegistriesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get registries params
func (o *GetRegistriesParams) WithName(name *string) *GetRegistriesParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get registries params
func (o *GetRegistriesParams) SetName(name *string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegistriesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}