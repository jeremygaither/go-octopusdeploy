// Code generated by go-swagger; DO NOT EDIT.

package teams

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

// NewLoadResponseDescriptorUsersTeamTeamResourceParams creates a new LoadResponseDescriptorUsersTeamTeamResourceParams object
// with the default values initialized.
func NewLoadResponseDescriptorUsersTeamTeamResourceParams() *LoadResponseDescriptorUsersTeamTeamResourceParams {
	var ()
	return &LoadResponseDescriptorUsersTeamTeamResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithTimeout creates a new LoadResponseDescriptorUsersTeamTeamResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithTimeout(timeout time.Duration) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	var ()
	return &LoadResponseDescriptorUsersTeamTeamResourceParams{

		timeout: timeout,
	}
}

// NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithContext creates a new LoadResponseDescriptorUsersTeamTeamResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithContext(ctx context.Context) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	var ()
	return &LoadResponseDescriptorUsersTeamTeamResourceParams{

		Context: ctx,
	}
}

// NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithHTTPClient creates a new LoadResponseDescriptorUsersTeamTeamResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadResponseDescriptorUsersTeamTeamResourceParamsWithHTTPClient(client *http.Client) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	var ()
	return &LoadResponseDescriptorUsersTeamTeamResourceParams{
		HTTPClient: client,
	}
}

/*LoadResponseDescriptorUsersTeamTeamResourceParams contains all the parameters to send to the API endpoint
for the load response descriptor users team team resource operation typically these are written to a http.Request
*/
type LoadResponseDescriptorUsersTeamTeamResourceParams struct {

	/*ID
	  ID of the TeamResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) WithTimeout(timeout time.Duration) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) WithContext(ctx context.Context) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) WithHTTPClient(client *http.Client) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) WithID(id string) *LoadResponseDescriptorUsersTeamTeamResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the load response descriptor users team team resource params
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LoadResponseDescriptorUsersTeamTeamResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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