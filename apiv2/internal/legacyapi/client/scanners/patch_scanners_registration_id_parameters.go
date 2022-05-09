// Code generated by go-swagger; DO NOT EDIT.

package scanners

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

	"github.com/zelat/goharbor-client/apiv2/model/legacy"
)

// NewPatchScannersRegistrationIDParams creates a new PatchScannersRegistrationIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchScannersRegistrationIDParams() *PatchScannersRegistrationIDParams {
	return &PatchScannersRegistrationIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchScannersRegistrationIDParamsWithTimeout creates a new PatchScannersRegistrationIDParams object
// with the ability to set a timeout on a request.
func NewPatchScannersRegistrationIDParamsWithTimeout(timeout time.Duration) *PatchScannersRegistrationIDParams {
	return &PatchScannersRegistrationIDParams{
		timeout: timeout,
	}
}

// NewPatchScannersRegistrationIDParamsWithContext creates a new PatchScannersRegistrationIDParams object
// with the ability to set a context for a request.
func NewPatchScannersRegistrationIDParamsWithContext(ctx context.Context) *PatchScannersRegistrationIDParams {
	return &PatchScannersRegistrationIDParams{
		Context: ctx,
	}
}

// NewPatchScannersRegistrationIDParamsWithHTTPClient creates a new PatchScannersRegistrationIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchScannersRegistrationIDParamsWithHTTPClient(client *http.Client) *PatchScannersRegistrationIDParams {
	return &PatchScannersRegistrationIDParams{
		HTTPClient: client,
	}
}

/* PatchScannersRegistrationIDParams contains all the parameters to send to the API endpoint
   for the patch scanners registration ID operation.

   Typically these are written to a http.Request.
*/
type PatchScannersRegistrationIDParams struct {

	// Payload.
	Payload *legacy.IsDefault

	/* RegistrationID.

	   The scanner registration identifier.
	*/
	RegistrationID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch scanners registration ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchScannersRegistrationIDParams) WithDefaults() *PatchScannersRegistrationIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch scanners registration ID params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchScannersRegistrationIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) WithTimeout(timeout time.Duration) *PatchScannersRegistrationIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) WithContext(ctx context.Context) *PatchScannersRegistrationIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) WithHTTPClient(client *http.Client) *PatchScannersRegistrationIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPayload adds the payload to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) WithPayload(payload *legacy.IsDefault) *PatchScannersRegistrationIDParams {
	o.SetPayload(payload)
	return o
}

// SetPayload adds the payload to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) SetPayload(payload *legacy.IsDefault) {
	o.Payload = payload
}

// WithRegistrationID adds the registrationID to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) WithRegistrationID(registrationID string) *PatchScannersRegistrationIDParams {
	o.SetRegistrationID(registrationID)
	return o
}

// SetRegistrationID adds the registrationId to the patch scanners registration ID params
func (o *PatchScannersRegistrationIDParams) SetRegistrationID(registrationID string) {
	o.RegistrationID = registrationID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchScannersRegistrationIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Payload != nil {
		if err := r.SetBodyParam(o.Payload); err != nil {
			return err
		}
	}

	// path param registration_id
	if err := r.SetPathParam("registration_id", o.RegistrationID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
