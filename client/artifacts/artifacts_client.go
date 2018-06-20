// Code generated by go-swagger; DO NOT EDIT.

package artifacts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new artifacts API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for artifacts API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateResponseDescriptorServerTasksArtifactArtifactResource creates a artifact resource

Creates a new artifacts.
*/
func (a *Client) CreateResponseDescriptorServerTasksArtifactArtifactResource(params *CreateResponseDescriptorServerTasksArtifactArtifactResourceParams, authInfo runtime.ClientAuthInfoWriter) (*CreateResponseDescriptorServerTasksArtifactArtifactResourceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateResponseDescriptorServerTasksArtifactArtifactResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateResponseDescriptor[ServerTasks.Artifact,ArtifactResource]",
		Method:             "POST",
		PathPattern:        "/api/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateResponseDescriptorServerTasksArtifactArtifactResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateResponseDescriptorServerTasksArtifactArtifactResourceCreated), nil

}

/*
CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponder Lists all of the artifacts in the current Octopus installation, from all releases. The results will be sorted by date from most recently to least recently created.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponder(params *CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponderParams, authInfo runtime.ClientAuthInfoWriter) (*CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponderOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponderParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CustomQueryResponseDescriptor[Octopus.Server.Web.Api.Actions.ListArtifactsResponder]",
		Method:             "GET",
		PathPattern:        "/api/artifacts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponderReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CustomQueryResponseDescriptorOctopusServerWebAPIActionsListArtifactsResponderOK), nil

}

/*
DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResource deletes a artifact resource by ID

Deletes an existing artifact.
*/
func (a *Client) DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResource(params *DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResourceParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteOnBackgroundResponseDescriptor[ServerTasks.Artifact,ArtifactResource]",
		Method:             "DELETE",
		PathPattern:        "/api/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteOnBackgroundResponseDescriptorServerTasksArtifactArtifactResourceOK), nil

}

/*
FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentAction Sets the content associated with an artifact.
NOTE: This definition is not complete. We will be adding more detail in future releases of Octopus.
*/
func (a *Client) FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentAction(params *FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionParams, authInfo runtime.ClientAuthInfoWriter, writer io.Writer) (*FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionOK, *FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewFileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "FileResponseDescriptor[Octopus.Server.Web.Api.Actions.ArtifactContentAction]",
		Method:             "PUT",
		PathPattern:        "/api/artifacts/{id}/content",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionReader{formats: a.formats, writer: writer},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionOK:
		return value, nil, nil
	case *FileResponseDescriptorOctopusServerWebAPIActionsArtifactContentActionNoContent:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
LoadResponseDescriptorServerTasksArtifactArtifactResource gets a artifact resource by ID

Gets a single artifact by ID.
*/
func (a *Client) LoadResponseDescriptorServerTasksArtifactArtifactResource(params *LoadResponseDescriptorServerTasksArtifactArtifactResourceParams, authInfo runtime.ClientAuthInfoWriter) (*LoadResponseDescriptorServerTasksArtifactArtifactResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadResponseDescriptorServerTasksArtifactArtifactResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "LoadResponseDescriptor[ServerTasks.Artifact,ArtifactResource]",
		Method:             "GET",
		PathPattern:        "/api/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadResponseDescriptorServerTasksArtifactArtifactResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadResponseDescriptorServerTasksArtifactArtifactResourceOK), nil

}

/*
ModifyResponseDescriptorServerTasksArtifactArtifactResource modifies a artifact resource by ID

Modifies an existing artifact.
*/
func (a *Client) ModifyResponseDescriptorServerTasksArtifactArtifactResource(params *ModifyResponseDescriptorServerTasksArtifactArtifactResourceParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyResponseDescriptorServerTasksArtifactArtifactResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyResponseDescriptorServerTasksArtifactArtifactResourceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ModifyResponseDescriptor[ServerTasks.Artifact,ArtifactResource]",
		Method:             "PUT",
		PathPattern:        "/api/artifacts/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ModifyResponseDescriptorServerTasksArtifactArtifactResourceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyResponseDescriptorServerTasksArtifactArtifactResourceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}