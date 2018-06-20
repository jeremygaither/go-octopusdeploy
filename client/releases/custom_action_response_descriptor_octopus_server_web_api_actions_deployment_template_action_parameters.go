// Code generated by go-swagger; DO NOT EDIT.

package releases

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions deployment template action operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions deployment template action params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentTemplateActionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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