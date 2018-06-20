// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions tenant variables get responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions tenant variables get responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTenantVariablesGetResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}