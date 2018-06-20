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

// DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader is a Reader for the DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResource structure.
type DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK creates a DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK with default headers values
func NewDeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK() *DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK {
	return &DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK{}
}

/*DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK handles this case with default header values.

EnvironmentResource Deleted
*/
type DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK struct {
	Payload *models.TaskResource
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/environments/{id}][%d] deleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK  %+v", 200, o.Payload)
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsDeploymentEnvironmentEnvironmentResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}