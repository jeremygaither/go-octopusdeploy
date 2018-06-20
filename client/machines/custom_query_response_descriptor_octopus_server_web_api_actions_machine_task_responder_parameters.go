// Code generated by go-swagger; DO NOT EDIT.

package machines

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

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams object
// with the default values initialized.
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams() *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	var ()
	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithTimeout creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	var ()
	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams{

		timeout: timeout,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithContext creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams object
// with the default values initialized, and the ability to set a context for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	var ()
	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams{

		Context: ctx,
	}
}

// NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithHTTPClient creates a new CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParamsWithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	var ()
	return &CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams{
		HTTPClient: client,
	}
}

/*CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams contains all the parameters to send to the API endpoint
for the custom query response descriptor octopus server web Api actions machine task responder operation typically these are written to a http.Request
*/
type CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams struct {

	/*ID
	  ID of the machine

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) WithTimeout(timeout time.Duration) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) WithContext(ctx context.Context) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) WithHTTPClient(client *http.Client) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) WithID(id string) *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the custom query response descriptor octopus server web Api actions machine task responder params
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CustomQueryResponseDescriptorOctopusServerWebAPIActionsMachineTaskResponderParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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