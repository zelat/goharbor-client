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
)

// NewDeleteProjectsProjectIDParams creates a new DeleteProjectsProjectIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteProjectsProjectIDParams() *DeleteProjectsProjectIDParams {
	return &DeleteProjectsProjectIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteProjectsProjectIDParamsWithTimeout creates a new DeleteProjectsProjectIDParams object
// with the ability to set a timeout on a request.
func NewDeleteProjectsProjectIDParamsWithTimeout(timeout time.Duration) *DeleteProjectsProjectIDParams {
	return &DeleteProjectsProjectIDParams{
		timeout: timeout,
	}
}

// NewDeleteProjectsProjectIDParamsWithContext creates a new DeleteProjectsProjectIDParams object
// with the ability to set a context for a request.
func NewDeleteProjectsProjectIDParamsWithContext(ctx context.Context) *DeleteProjectsProjectIDParams {
	return &DeleteProjectsProjectIDParams{
		Context: ctx,
	}
}

// NewDeleteProjectsProjectIDParamsWithHTTPClient creates a new DeleteProjectsProjectIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteProjectsProjectIDParamsWithHTTPClient(client *http.Client) *DeleteProjectsProjectIDParams {
	return &DeleteProjectsProjectIDParams{
		HTTPClient: client,
	}
}

/* DeleteProjectsProjectIDParams contains all the parameters to send to the API endpoint
   for the delete projects project ID operation.

   Typically these are written to a http.Request.
*/
type DeleteProjectsProjectIDParams struct {

	/* ProjectID.

	   Project ID of project which will be deleted.

	   Format: int64
	*/
	ProjectID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectsProjectIDParams) WithDefaults() *DeleteProjectsProjectIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete projects project ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteProjectsProjectIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) WithTimeout(timeout time.Duration) *DeleteProjectsProjectIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) WithContext(ctx context.Context) *DeleteProjectsProjectIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) WithHTTPClient(client *http.Client) *DeleteProjectsProjectIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) WithProjectID(projectID int64) *DeleteProjectsProjectIDParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the delete projects project ID params
func (o *DeleteProjectsProjectIDParams) SetProjectID(projectID int64) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteProjectsProjectIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param project_id
	if err := r.SetPathParam("project_id", swag.FormatInt64(o.ProjectID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}