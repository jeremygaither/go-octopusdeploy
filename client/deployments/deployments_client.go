// Code generated by go-swagger; DO NOT EDIT.

package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new deployments API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for deployments API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateAction Creates a new deployment.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateAction(params *CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionParams, authInfo runtime.ClientAuthInfoWriter) (*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomActionResponseDescriptor[Octopus.Server.Web.Api.Actions.DeploymentCreateAction]",
		Method:             "POST",
		PathPattern:        "/api/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomActionResponseDescriptorOctopusServerWebAPIActionsDeploymentCreateActionCreated), nil

}

/*
CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponder Lists all of the deployments in the current Octopus installation, from projects, releases and environments accessible by the current user. The results will be sorted from most recent to least recent deployment.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponder(params *CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomQueryResponseDescriptor[Octopus.Server.Web.Api.Actions.DeploymentsQueryResponder]",
		Method:             "GET",
		PathPattern:        "/api/deployments",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomQueryResponseDescriptorOctopusServerWebAPIActionsDeploymentsQueryResponderOK), nil

}

/*
DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResource deletes a deployment resource by ID

Deletes a deployment.
*/
func (a *Client) DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResource(params *DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResourceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOnBackgroundResponseDescriptor[Projects.Deployment,DeploymentResource]",
		Method:             "DELETE",
		PathPattern:        "/api/deployments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOnBackgroundResponseDescriptorProjectsDeploymentDeploymentResourceOK), nil

}

/*
LoadResponseDescriptorProjectsDeploymentDeploymentResource gets a deployment resource by ID

Gets a deployment by ID.
*/
func (a *Client) LoadResponseDescriptorProjectsDeploymentDeploymentResource(params *LoadResponseDescriptorProjectsDeploymentDeploymentResourceParams, authInfo runtime.ClientAuthInfoWriter) (*LoadResponseDescriptorProjectsDeploymentDeploymentResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadResponseDescriptorProjectsDeploymentDeploymentResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoadResponseDescriptor[Projects.Deployment,DeploymentResource]",
		Method:             "GET",
		PathPattern:        "/api/deployments/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadResponseDescriptorProjectsDeploymentDeploymentResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadResponseDescriptorProjectsDeploymentDeploymentResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}