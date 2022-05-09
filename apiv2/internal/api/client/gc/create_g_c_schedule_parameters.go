// Code generated by go-swagger; DO NOT EDIT.

package gc

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

	"github.com/zelat/goharbor-client/apiv2/model"
)

// NewCreateGCScheduleParams creates a new CreateGCScheduleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateGCScheduleParams() *CreateGCScheduleParams {
	return &CreateGCScheduleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateGCScheduleParamsWithTimeout creates a new CreateGCScheduleParams object
// with the ability to set a timeout on a request.
func NewCreateGCScheduleParamsWithTimeout(timeout time.Duration) *CreateGCScheduleParams {
	return &CreateGCScheduleParams{
		timeout: timeout,
	}
}

// NewCreateGCScheduleParamsWithContext creates a new CreateGCScheduleParams object
// with the ability to set a context for a request.
func NewCreateGCScheduleParamsWithContext(ctx context.Context) *CreateGCScheduleParams {
	return &CreateGCScheduleParams{
		Context: ctx,
	}
}

// NewCreateGCScheduleParamsWithHTTPClient creates a new CreateGCScheduleParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateGCScheduleParamsWithHTTPClient(client *http.Client) *CreateGCScheduleParams {
	return &CreateGCScheduleParams{
		HTTPClient: client,
	}
}

/* CreateGCScheduleParams contains all the parameters to send to the API endpoint
   for the create g c schedule operation.

   Typically these are written to a http.Request.
*/
type CreateGCScheduleParams struct {

	/* Schedule.

	   Updates of gc's schedule.
	*/
	Schedule *model.Schedule

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create g c schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGCScheduleParams) WithDefaults() *CreateGCScheduleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create g c schedule params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateGCScheduleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create g c schedule params
func (o *CreateGCScheduleParams) WithTimeout(timeout time.Duration) *CreateGCScheduleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create g c schedule params
func (o *CreateGCScheduleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create g c schedule params
func (o *CreateGCScheduleParams) WithContext(ctx context.Context) *CreateGCScheduleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create g c schedule params
func (o *CreateGCScheduleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create g c schedule params
func (o *CreateGCScheduleParams) WithHTTPClient(client *http.Client) *CreateGCScheduleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create g c schedule params
func (o *CreateGCScheduleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSchedule adds the schedule to the create g c schedule params
func (o *CreateGCScheduleParams) WithSchedule(schedule *model.Schedule) *CreateGCScheduleParams {
	o.SetSchedule(schedule)
	return o
}

// SetSchedule adds the schedule to the create g c schedule params
func (o *CreateGCScheduleParams) SetSchedule(schedule *model.Schedule) {
	o.Schedule = schedule
}

// WriteToRequest writes these params to a swagger request
func (o *CreateGCScheduleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Schedule != nil {
		if err := r.SetBodyParam(o.Schedule); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
