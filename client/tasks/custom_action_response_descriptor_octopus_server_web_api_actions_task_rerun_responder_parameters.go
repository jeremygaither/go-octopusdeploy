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

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams object
// with the default values initialized.
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithTimeout creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams{

		timeout: timeout,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithContext creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams{

		Context: ctx,
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithHTTPClient creates a new CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParamsWithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	var ()
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams{
		HTTPClient: client,
	}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams contains all the parameters to send to the API endpoint
for the custom action response descriptor octopus server web Api actions task rerun responder operation typically these are written to a http.Request
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams struct {

	/*ID
	  ID of the resource

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) WithTimeout(timeout time.Duration) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) WithContext(ctx context.Context) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) WithHTTPClient(client *http.Client) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) WithID(id string) *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom action response descriptor octopus server web Api actions task rerun responder params
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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