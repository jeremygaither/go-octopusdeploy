// Code generated by go-swagger; DO NOT EDIT.

package project_groups

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

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams object
// with the default values initialized.
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams() *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithTimeout creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams{

		timeout: timeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithContext creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams{

		Context: ctx,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithHTTPClient creates a new DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParamsWithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams{
		HTTPClient: client,
	}
}

/*DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams contains all the parameters to send to the API endpoint
for the delete on background response descriptor projects project group project group resource operation typically these are written to a http.Request
*/
type DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams struct {

	/*ID
	  ID of the ProjectGroupResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) WithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) WithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) WithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) WithID(id string) *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete on background response descriptor projects project group project group resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOnBackgroundResponseDescriptorProjectsProjectGroupProjectGroupResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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