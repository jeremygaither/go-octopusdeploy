// Code generated by go-swagger; DO NOT EDIT.

package upgrade_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new upgrade configuration API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for upgrade configuration API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetAction Gets information about the upgrade configuration in use by the Octopus Server.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.UpgradeConfigurationGetAction]",
		Method:             "GET",
		PathPattern:        "/api/upgradeconfiguration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationGetActionOK), nil

}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateAction Updates the upgrade configuration used by the Octopus Server.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.UpgradeConfigurationUpdateAction]",
		Method:             "PUT",
		PathPattern:        "/api/upgradeconfiguration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsUpgradeConfigurationUpdateActionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}