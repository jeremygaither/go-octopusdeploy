// Code generated by go-swagger; DO NOT EDIT.

package upgrade_configuration

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions upgrade configuration get action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions upgrade configuration get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}