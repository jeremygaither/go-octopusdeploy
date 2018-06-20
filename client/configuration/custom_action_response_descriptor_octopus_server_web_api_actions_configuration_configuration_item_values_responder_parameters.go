// Code generated by go-swagger; DO NOT EDIT.

package configuration

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions configuration configuration item values responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions configuration configuration item values responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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