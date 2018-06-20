// Code generated by go-swagger; DO NOT EDIT.

package environments

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

// NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams creates a new LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams object
// with the default values initialized.
func NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams() *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	var ()
	return &LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithTimeout creates a new LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithTimeout(timeout time.Duration) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	var ()
	return &LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams{

		timeout: timeout,
	}
}

// NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithContext creates a new LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithContext(ctx context.Context) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	var ()
	return &LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams{

		Context: ctx,
	}
}

// NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithHTTPClient creates a new LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParamsWithHTTPClient(client *http.Client) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	var ()
	return &LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams{
		HTTPClient: client,
	}
}

/*LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams contains all the parameters to send to the API endpoint
for the load response descriptor environments deployment environment environment resource operation typically these are written to a http.Request
*/
type LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams struct {

	/*ID
	  ID of the EnvironmentResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) WithTimeout(timeout time.Duration) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) WithContext(ctx context.Context) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) WithHTTPClient(client *http.Client) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) WithID(id string) *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the load response descriptor environments deployment environment environment resource params
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LoadResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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