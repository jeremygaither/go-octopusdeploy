// Code generated by go-swagger; DO NOT EDIT.

package tasks

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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions task state responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions task state responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskStateResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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