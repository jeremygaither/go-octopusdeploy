// Code generated by go-swagger; DO NOT EDIT.

package octopus_server_nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams creates a new IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized.
func NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams() *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	var ()
	return &IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithTimeout creates a new IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithTimeout(timeout time.Duration) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	var ()
	return &IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		timeout: timeout,
	}
}

// NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithContext creates a new IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a context for a request
func NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithContext(ctx context.Context) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	var ()
	return &IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{

		Context: ctx,
	}
}

// NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithHTTPClient creates a new IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParamsWithHTTPClient(client *http.Client) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	var ()
	return &IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams{
		HTTPClient: client,
	}
}

/*IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams contains all the parameters to send to the API endpoint
for the index response descriptor clustering octopus server node octopus server node resource operation typically these are written to a http.Request
*/
type IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams struct {

	/*Skip
	  Number of items to skip

	*/
	Skip *int32
	/*Take
	  Number of items to take

	*/
	Take *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithTimeout(timeout time.Duration) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithContext(ctx context.Context) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithHTTPClient(client *http.Client) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSkip adds the skip to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithSkip(skip *int32) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetSkip(skip)
	return o
}

// SetSkip adds the skip to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetSkip(skip *int32) {
	o.Skip = skip
}

// WithTake adds the take to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WithTake(take *int32) *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams {
	o.SetTake(take)
	return o
}

// SetTake adds the take to the index response descriptor clustering octopus server node octopus server node resource params
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) SetTake(take *int32) {
	o.Take = take
}

// WriteToRequest writes these params to a swagger request
func (o *IndexResponseDescriptorClusteringOctopusServerNodeOctopusServerNodeResourceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Skip != nil {

		// query param skip
		var qrSkip int32
		if o.Skip != nil {
			qrSkip = *o.Skip
		}
		qSkip := swag.FormatInt32(qrSkip)
		if qSkip != "" {
			if err := r.SetQueryParam("skip", qSkip); err != nil {
				return err
			}
		}

	}

	if o.Take != nil {

		// query param take
		var qrTake int32
		if o.Take != nil {
			qrTake = *o.Take
		}
		qTake := swag.FormatInt32(qrTake)
		if qTake != "" {
			if err := r.SetQueryParam("take", qTake); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}