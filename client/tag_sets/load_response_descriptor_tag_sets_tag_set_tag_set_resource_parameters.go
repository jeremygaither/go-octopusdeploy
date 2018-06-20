// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

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

// NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParams creates a new LoadResponseDescriptorTagSetsTagSetTagSetResourceParams object
// with the default values initialized.
func NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParams() *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	var ()
	return &LoadResponseDescriptorTagSetsTagSetTagSetResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithTimeout creates a new LoadResponseDescriptorTagSetsTagSetTagSetResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithTimeout(timeout time.Duration) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	var ()
	return &LoadResponseDescriptorTagSetsTagSetTagSetResourceParams{

		timeout: timeout,
	}
}

// NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithContext creates a new LoadResponseDescriptorTagSetsTagSetTagSetResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithContext(ctx context.Context) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	var ()
	return &LoadResponseDescriptorTagSetsTagSetTagSetResourceParams{

		Context: ctx,
	}
}

// NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithHTTPClient creates a new LoadResponseDescriptorTagSetsTagSetTagSetResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoadResponseDescriptorTagSetsTagSetTagSetResourceParamsWithHTTPClient(client *http.Client) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	var ()
	return &LoadResponseDescriptorTagSetsTagSetTagSetResourceParams{
		HTTPClient: client,
	}
}

/*LoadResponseDescriptorTagSetsTagSetTagSetResourceParams contains all the parameters to send to the API endpoint
for the load response descriptor tag sets tag set tag set resource operation typically these are written to a http.Request
*/
type LoadResponseDescriptorTagSetsTagSetTagSetResourceParams struct {

	/*ID
	  ID of the TagSetResource to load

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) WithTimeout(timeout time.Duration) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) WithContext(ctx context.Context) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) WithHTTPClient(client *http.Client) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) WithID(id string) *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the load response descriptor tag sets tag set tag set resource params
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *LoadResponseDescriptorTagSetsTagSetTagSetResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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