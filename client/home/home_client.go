// Code generated by go-swagger; DO NOT EDIT.

package home

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new home API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for home API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponder Returns a document describing the current Octopus Server and links to other parts of the API.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponderParams) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.HomeResponder]",
		Method:             "GET",
		PathPattern:        "/api",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponderReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsHomeResponderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}