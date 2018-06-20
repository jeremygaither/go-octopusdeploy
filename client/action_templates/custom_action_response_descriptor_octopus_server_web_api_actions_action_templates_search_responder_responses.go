// Code generated by go-swagger; DO NOT EDIT.

package action_templates

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK handles this case with default header values.

IEnumerable[ActionTemplateSearchResource] resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK struct {
	Payload []*models.ActionTemplateSearchResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/actiontemplates/search][%d] customActionResponseDescriptorOctopusServerWebApiActionsActionTemplatesSearchResponderOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsActionTemplatesSearchResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}