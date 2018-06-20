// Code generated by go-swagger; DO NOT EDIT.

package machine_policies

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

// NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams creates a new CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams object
// with the default values initialized.
func NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams() *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	var ()
	return &CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithTimeout creates a new CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithTimeout(timeout time.Duration) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	var ()
	return &CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams{

		timeout: timeout,
	}
}

// NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithContext creates a new CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithContext(ctx context.Context) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	var ()
	return &CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams{

		Context: ctx,
	}
}

// NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithHTTPClient creates a new CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParamsWithHTTPClient(client *http.Client) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	var ()
	return &CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams{
		HTTPClient: client,
	}
}

/*CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams contains all the parameters to send to the API endpoint
for the create response descriptor policies machine policy machine policy resource operation typically these are written to a http.Request
*/
type CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams struct {

	/*MachinePolicyResource
	  The MachinePolicyResource resource to create

	*/
	MachinePolicyResource *models.MachinePolicyResource

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) WithTimeout(timeout time.Duration) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) WithContext(ctx context.Context) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) WithHTTPClient(client *http.Client) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMachinePolicyResource adds the machinePolicyResource to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) WithMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams {
	o.SetMachinePolicyResource(machinePolicyResource)
	return o
}

// SetMachinePolicyResource adds the machinePolicyResource to the create response descriptor policies machine policy machine policy resource params
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) SetMachinePolicyResource(machinePolicyResource *models.MachinePolicyResource) {
	o.MachinePolicyResource = machinePolicyResource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResponseDescriptorPoliciesMachinePolicyMachinePolicyResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.MachinePolicyResource != nil {
		if err := r.SetBodyParam(o.MachinePolicyResource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}