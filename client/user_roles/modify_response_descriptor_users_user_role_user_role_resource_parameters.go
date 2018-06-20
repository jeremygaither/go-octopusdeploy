// Code generated by go-swagger; DO NOT EDIT.

package user_roles

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

// NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParams creates a new ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams object
// with the default values initialized.
func NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParams() *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	var ()
	return &ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithTimeout creates a new ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithTimeout(timeout time.Duration) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	var ()
	return &ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams{

		timeout: timeout,
	}
}

// NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithContext creates a new ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithContext(ctx context.Context) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	var ()
	return &ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams{

		Context: ctx,
	}
}

// NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithHTTPClient creates a new ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewModifyResponseDescriptorUsersUserRoleUserRoleResourceParamsWithHTTPClient(client *http.Client) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	var ()
	return &ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams{
		HTTPClient: client,
	}
}

/*ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams contains all the parameters to send to the API endpoint
for the modify response descriptor users user role user role resource operation typically these are written to a http.Request
*/
type ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams struct {

	/*UserRoleResource
	  The UserRoleResource resource to create

	*/
	UserRoleResource *models.UserRoleResource
	/*ID
	  ID of the UserRoleResource to modify

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WithTimeout(timeout time.Duration) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WithContext(ctx context.Context) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WithHTTPClient(client *http.Client) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUserRoleResource adds the userRoleResource to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WithUserRoleResource(userRoleResource *models.UserRoleResource) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	o.SetUserRoleResource(userRoleResource)
	return o
}

// SetUserRoleResource adds the userRoleResource to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) SetUserRoleResource(userRoleResource *models.UserRoleResource) {
	o.UserRoleResource = userRoleResource
}

// WithID adds the id to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WithID(id string) *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the modify response descriptor users user role user role resource params
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ModifyResponseDescriptorUsersUserRoleUserRoleResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.UserRoleResource != nil {
		if err := r.SetBodyParam(o.UserRoleResource); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}