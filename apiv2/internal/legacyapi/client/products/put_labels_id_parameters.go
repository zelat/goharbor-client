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
	"github.com/go-openapi/swag"

	"github.com/zelat/goharbor-client/v4/apiv2/model/legacy"
)

// NewPutLabelsIDParams creates a new PutLabelsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPutLabelsIDParams() *PutLabelsIDParams {
	return &PutLabelsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPutLabelsIDParamsWithTimeout creates a new PutLabelsIDParams object
// with the ability to set a timeout on a request.
func NewPutLabelsIDParamsWithTimeout(timeout time.Duration) *PutLabelsIDParams {
	return &PutLabelsIDParams{
		timeout: timeout,
	}
}

// NewPutLabelsIDParamsWithContext creates a new PutLabelsIDParams object
// with the ability to set a context for a request.
func NewPutLabelsIDParamsWithContext(ctx context.Context) *PutLabelsIDParams {
	return &PutLabelsIDParams{
		Context: ctx,
	}
}

// NewPutLabelsIDParamsWithHTTPClient creates a new PutLabelsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPutLabelsIDParamsWithHTTPClient(client *http.Client) *PutLabelsIDParams {
	return &PutLabelsIDParams{
		HTTPClient: client,
	}
}

/* PutLabelsIDParams contains all the parameters to send to the API endpoint
   for the put labels ID operation.

   Typically these are written to a http.Request.
*/
type PutLabelsIDParams struct {

	/* ID.

	   Label ID

	   Format: int64
	*/
	ID int64

	/* Label.

	   The updated label json object.
	*/
	Label *legacy.Label

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the put labels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLabelsIDParams) WithDefaults() *PutLabelsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the put labels ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PutLabelsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the put labels ID params
func (o *PutLabelsIDParams) WithTimeout(timeout time.Duration) *PutLabelsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put labels ID params
func (o *PutLabelsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put labels ID params
func (o *PutLabelsIDParams) WithContext(ctx context.Context) *PutLabelsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put labels ID params
func (o *PutLabelsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put labels ID params
func (o *PutLabelsIDParams) WithHTTPClient(client *http.Client) *PutLabelsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put labels ID params
func (o *PutLabelsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the put labels ID params
func (o *PutLabelsIDParams) WithID(id int64) *PutLabelsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put labels ID params
func (o *PutLabelsIDParams) SetID(id int64) {
	o.ID = id
}

// WithLabel adds the label to the put labels ID params
func (o *PutLabelsIDParams) WithLabel(label *legacy.Label) *PutLabelsIDParams {
	o.SetLabel(label)
	return o
}

// SetLabel adds the label to the put labels ID params
func (o *PutLabelsIDParams) SetLabel(label *legacy.Label) {
	o.Label = label
}

// WriteToRequest writes these params to a swagger request
func (o *PutLabelsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}
	if o.Label != nil {
		if err := r.SetBodyParam(o.Label); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
