// Code generated by go-swagger; DO NOT EDIT.

package users

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {

	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions users user identity metadata action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions users user identity metadata action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserIdentityMetadataActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}