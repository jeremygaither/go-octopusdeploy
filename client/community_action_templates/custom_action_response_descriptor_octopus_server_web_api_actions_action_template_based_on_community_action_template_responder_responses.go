// Code generated by go-swagger; DO NOT EDIT.

package community_action_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK handles this case with default header values.

ActionTemplateResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK struct {
	Payload *models.ActionTemplateResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/communityactiontemplates/{id}/actiontemplate][%d] customActionResponseDescriptorOctopusServerWebApiActionsActionTemplateBasedOnCommunityActionTemplateResponderOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ActionTemplateResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/communityactiontemplates/{id}/actiontemplate][%d] customActionResponseDescriptorOctopusServerWebApiActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplateBasedOnCommunityActionTemplateResponderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}