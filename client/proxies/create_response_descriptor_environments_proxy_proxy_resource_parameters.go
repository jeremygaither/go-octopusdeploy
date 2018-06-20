// Code generated by go-swagger; DO NOT EDIT.

package proxies

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

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParams creates a new CreateResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized.
func NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParams() *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	var ()
	return &CreateResponseDescriptorEnvironmentsProxyProxyResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithTimeout creates a new CreateResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithTimeout(timeout time.Duration) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	var ()
	return &CreateResponseDescriptorEnvironmentsProxyProxyResourceParams{

		timeout: timeout,
	}
}

// NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithContext creates a new CreateResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithContext(ctx context.Context) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	var ()
	return &CreateResponseDescriptorEnvironmentsProxyProxyResourceParams{

		Context: ctx,
	}
}

// NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithHTTPClient creates a new CreateResponseDescriptorEnvironmentsProxyProxyResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateResponseDescriptorEnvironmentsProxyProxyResourceParamsWithHTTPClient(client *http.Client) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	var ()
	return &CreateResponseDescriptorEnvironmentsProxyProxyResourceParams{
		HTTPClient: client,
	}
}

/*CreateResponseDescriptorEnvironmentsProxyProxyResourceParams contains all the parameters to send to the API endpoint
for the create response descriptor environments proxy proxy resource operation typically these are written to a http.Request
*/
type CreateResponseDescriptorEnvironmentsProxyProxyResourceParams struct {

	/*ProxyResource
	  The ProxyResource resource to create

	*/
	ProxyResource *models.ProxyResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) WithTimeout(timeout time.Duration) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) WithContext(ctx context.Context) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) WithHTTPClient(client *http.Client) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProxyResource adds the proxyResource to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) WithProxyResource(proxyResource *models.ProxyResource) *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams {
	o.SetProxyResource(proxyResource)
	return o
}

// SetProxyResource adds the proxyResource to the create response descriptor environments proxy proxy resource params
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) SetProxyResource(proxyResource *models.ProxyResource) {
	o.ProxyResource = proxyResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ProxyResource != nil {
		if err := r.SetBodyParam(o.ProxyResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}