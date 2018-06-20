// Code generated by go-swagger; DO NOT EDIT.

package subscription

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

// NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams creates a new ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams object
// with the default values initialized.
func NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams() *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {

	return &ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithTimeout creates a new ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithTimeout(timeout time.Duration) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {

	return &ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams{

		timeout: timeout,
	}
}

// NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithContext creates a new ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithContext(ctx context.Context) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {

	return &ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams{

		Context: ctx,
	}
}

// NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithHTTPClient creates a new ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParamsWithHTTPClient(client *http.Client) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {

	return &ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams{
		HTTPClient: client,
	}
}

/*ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams contains all the parameters to send to the API endpoint
for the list all response descriptor subscriptions subscription subscription resource operation typically these are written to a http.Request
*/
type ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) WithTimeout(timeout time.Duration) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) WithContext(ctx context.Context) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) WithHTTPClient(client *http.Client) *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list all response descriptor subscriptions subscription subscription resource params
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ListAllResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}