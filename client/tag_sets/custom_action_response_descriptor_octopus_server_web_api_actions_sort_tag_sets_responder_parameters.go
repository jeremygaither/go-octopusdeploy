// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions sort tag sets responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions sort tag sets responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}