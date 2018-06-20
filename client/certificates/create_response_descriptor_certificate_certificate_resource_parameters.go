// Code generated by go-swagger; DO NOT EDIT.

package certificates

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

// NewCreateResponseDescriptorCertificateCertificateResourceParams creates a new CreateResponseDescriptorCertificateCertificateResourceParams object
// with the default values initialized.
func NewCreateResponseDescriptorCertificateCertificateResourceParams() *CreateResponseDescriptorCertificateCertificateResourceParams {
	var ()
	return &CreateResponseDescriptorCertificateCertificateResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResponseDescriptorCertificateCertificateResourceParamsWithTimeout creates a new CreateResponseDescriptorCertificateCertificateResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateResponseDescriptorCertificateCertificateResourceParamsWithTimeout(timeout time.Duration) *CreateResponseDescriptorCertificateCertificateResourceParams {
	var ()
	return &CreateResponseDescriptorCertificateCertificateResourceParams{

		timeout: timeout,
	}
}

// NewCreateResponseDescriptorCertificateCertificateResourceParamsWithContext creates a new CreateResponseDescriptorCertificateCertificateResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateResponseDescriptorCertificateCertificateResourceParamsWithContext(ctx context.Context) *CreateResponseDescriptorCertificateCertificateResourceParams {
	var ()
	return &CreateResponseDescriptorCertificateCertificateResourceParams{

		Context: ctx,
	}
}

// NewCreateResponseDescriptorCertificateCertificateResourceParamsWithHTTPClient creates a new CreateResponseDescriptorCertificateCertificateResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateResponseDescriptorCertificateCertificateResourceParamsWithHTTPClient(client *http.Client) *CreateResponseDescriptorCertificateCertificateResourceParams {
	var ()
	return &CreateResponseDescriptorCertificateCertificateResourceParams{
		HTTPClient: client,
	}
}

/*CreateResponseDescriptorCertificateCertificateResourceParams contains all the parameters to send to the API endpoint
for the create response descriptor certificate certificate resource operation typically these are written to a http.Request
*/
type CreateResponseDescriptorCertificateCertificateResourceParams struct {

	/*CertificateResource
	  The CertificateResource resource to create

	*/
	CertificateResource *models.CertificateResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) WithTimeout(timeout time.Duration) *CreateResponseDescriptorCertificateCertificateResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) WithContext(ctx context.Context) *CreateResponseDescriptorCertificateCertificateResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) WithHTTPClient(client *http.Client) *CreateResponseDescriptorCertificateCertificateResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateResource adds the certificateResource to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) WithCertificateResource(certificateResource *models.CertificateResource) *CreateResponseDescriptorCertificateCertificateResourceParams {
	o.SetCertificateResource(certificateResource)
	return o
}

// SetCertificateResource adds the certificateResource to the create response descriptor certificate certificate resource params
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) SetCertificateResource(certificateResource *models.CertificateResource) {
	o.CertificateResource = certificateResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResponseDescriptorCertificateCertificateResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateResource != nil {
		if err := r.SetBodyParam(o.CertificateResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}