// Code generated by go-swagger; DO NOT EDIT.

package environments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader is a Reader for the CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResource structure.
type CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated creates a CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated with default headers values
func NewCreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated() *CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated {
	return &CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated{}
}

/*CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated handles this case with default header values.

EnvironmentResource Created
*/
type CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated struct {
	Payload *models.EnvironmentResource
}

func (o *CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated) Error() string {
	return fmt.Sprintf("[POST /api/environments][%d] createResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated  %+v", 201, o.Payload)
}

func (o *CreateResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EnvironmentResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}