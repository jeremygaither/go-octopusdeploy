// Code generated by go-swagger; DO NOT EDIT.

package machines

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceReader is a Reader for the DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResource structure.
type DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK creates a DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK with default headers values
func NewDeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK() *DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK {
	return &DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK{}
}

/*DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK handles this case with default header values.

MachineResource Deleted
*/
type DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK struct {
	Payload *models.TaskResource
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK) Error() string {
	return fmt.Sprintf("[DELETE /api/machines/{id}][%d] deleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK  %+v", 200, o.Payload)
}

func (o *DeleteOnBackgroundResponseDescriptorEnvironmentsMachineMachineResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}