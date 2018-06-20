// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK handles this case with default header values.

UserPermissionSetResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK struct {
	Payload *models.UserPermissionSetResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}/permissions][%d] customActionResponseDescriptorOctopusServerWebApiActionsUsersUserGetPermissionsActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.UserPermissionSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/users/{id}/permissions][%d] customActionResponseDescriptorOctopusServerWebApiActionsUsersUserGetPermissionsActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsUsersUserGetPermissionsActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}