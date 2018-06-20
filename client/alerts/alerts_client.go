// Code generated by go-swagger; DO NOT EDIT.

package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new alerts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alerts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponder Shows information about the status of the Octopus Server.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponder(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.AlertResponder]",
		Method:             "GET",
		PathPattern:        "/api/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsAlertResponderOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}