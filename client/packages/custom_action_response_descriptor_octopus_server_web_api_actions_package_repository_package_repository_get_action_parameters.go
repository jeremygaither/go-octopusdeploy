// Code generated by go-swagger; DO NOT EDIT.

package packages

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions package repository package repository get action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions package repository package repository get action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsPackageRepositoryPackageRepositoryGetActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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