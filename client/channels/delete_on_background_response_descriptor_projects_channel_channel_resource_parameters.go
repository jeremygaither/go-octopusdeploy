// Code generated by go-swagger; DO NOT EDIT.

package channels

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

// NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams creates a new DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams object
// with the default values initialized.
func NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams() *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithTimeout creates a new DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams{

		timeout: timeout,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithContext creates a new DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams{

		Context: ctx,
	}
}

// NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithHTTPClient creates a new DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParamsWithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	var ()
	return &DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams{
		HTTPClient: client,
	}
}

/*DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams contains all the parameters to send to the API endpoint
for the delete on background response descriptor projects channel channel resource operation typically these are written to a http.Request
*/
type DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams struct {

	/*ID
	  ID of the ChannelResource to delete

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) WithTimeout(timeout time.Duration) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) WithContext(ctx context.Context) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) WithHTTPClient(client *http.Client) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) WithID(id string) *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete on background response descriptor projects channel channel resource params
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOnBackgroundResponseDescriptorProjectsChannelChannelResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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