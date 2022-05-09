// Code generated by go-swagger; DO NOT EDIT.

package project

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

	"github.com/zelat/goharbor-client/apiv2/model"
)

// NewCreateProjectParams creates a new CreateProjectParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateProjectParams() *CreateProjectParams {
	return &CreateProjectParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectParamsWithTimeout creates a new CreateProjectParams object
// with the ability to set a timeout on a request.
func NewCreateProjectParamsWithTimeout(timeout time.Duration) *CreateProjectParams {
	return &CreateProjectParams{
		timeout: timeout,
	}
}

// NewCreateProjectParamsWithContext creates a new CreateProjectParams object
// with the ability to set a context for a request.
func NewCreateProjectParamsWithContext(ctx context.Context) *CreateProjectParams {
	return &CreateProjectParams{
		Context: ctx,
	}
}

// NewCreateProjectParamsWithHTTPClient creates a new CreateProjectParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateProjectParamsWithHTTPClient(client *http.Client) *CreateProjectParams {
	return &CreateProjectParams{
		HTTPClient: client,
	}
}

/* CreateProjectParams contains all the parameters to send to the API endpoint
   for the create project operation.

   Typically these are written to a http.Request.
*/
type CreateProjectParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string

	/* XResourceNameInLocation.

	   The flag to indicate whether to return the name of the resource in Location. When X-Resource-Name-In-Location is true, the Location will return the name of the resource.
	*/
	XResourceNameInLocation *bool

	/* Project.

	   New created project.
	*/
	Project *model.ProjectReq

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectParams) WithDefaults() *CreateProjectParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create project params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateProjectParams) SetDefaults() {
	var (
		xResourceNameInLocationDefault = bool(false)
	)

	val := CreateProjectParams{
		XResourceNameInLocation: &xResourceNameInLocationDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the create project params
func (o *CreateProjectParams) WithTimeout(timeout time.Duration) *CreateProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project params
func (o *CreateProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project params
func (o *CreateProjectParams) WithContext(ctx context.Context) *CreateProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project params
func (o *CreateProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project params
func (o *CreateProjectParams) WithHTTPClient(client *http.Client) *CreateProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project params
func (o *CreateProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the create project params
func (o *CreateProjectParams) WithXRequestID(xRequestID *string) *CreateProjectParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create project params
func (o *CreateProjectParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithXResourceNameInLocation adds the xResourceNameInLocation to the create project params
func (o *CreateProjectParams) WithXResourceNameInLocation(xResourceNameInLocation *bool) *CreateProjectParams {
	o.SetXResourceNameInLocation(xResourceNameInLocation)
	return o
}

// SetXResourceNameInLocation adds the xResourceNameInLocation to the create project params
func (o *CreateProjectParams) SetXResourceNameInLocation(xResourceNameInLocation *bool) {
	o.XResourceNameInLocation = xResourceNameInLocation
}

// WithProject adds the project to the create project params
func (o *CreateProjectParams) WithProject(project *model.ProjectReq) *CreateProjectParams {
	o.SetProject(project)
	return o
}

// SetProject adds the project to the create project params
func (o *CreateProjectParams) SetProject(project *model.ProjectReq) {
	o.Project = project
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.XResourceNameInLocation != nil {

		// header param X-Resource-Name-In-Location
		if err := r.SetHeaderParam("X-Resource-Name-In-Location", swag.FormatBool(*o.XResourceNameInLocation)); err != nil {
			return err
		}
	}
	if o.Project != nil {
		if err := r.SetBodyParam(o.Project); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
