// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK handles this case with default header values.

TaskResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK struct {
	Payload *models.TaskResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK) Error() string {
	return fmt.Sprintf("[POST /api/tasks/rerun/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsTaskRerunResponderOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.TaskResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest handles this case with default header values.

This task cannot be re-run.
No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest) Error() string {
	return fmt.Sprintf("[POST /api/tasks/rerun/{id}][%d] customActionResponseDescriptorOctopusServerWebApiActionsTaskRerunResponderBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsTaskRerunResponderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}