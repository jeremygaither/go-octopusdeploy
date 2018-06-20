// Code generated by go-swagger; DO NOT EDIT.

package projects

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

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams object
// with the default values initialized.
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams() *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithTimeout creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams{

		timeout: timeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithContext creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams{

		Context: ctx,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithHTTPClient creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParamsWithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams{
		HTTPClient: client,
	}
}

/*DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams contains all the parameters to send to the API endpoint
for the delete on background response descriptor projects project project resource operation typically these are written to a http.Request
*/
type DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams struct {

	/*ID
	  ID of the ProjectResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) WithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) WithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) WithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) WithID(id string) *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete on background response descriptor projects project project resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectProjectResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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