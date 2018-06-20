// Code generated by go-swagger; DO NOT EDIT.

package project_triggers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader is a Reader for the ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResource structure.
type ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK creates a ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK with default headers values
func NewModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK() *ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK {
	return &ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK{}
}

/*ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK handles this case with default header values.

ProjectTriggerResource Modified
*/
type ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK struct {
	Payload *models.ProjectTriggerResource
}

func (o *ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/projecttriggers/{id}][%d] modifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK  %+v", 200, o.Payload)
}

func (o *ModifyResponseDescriptorProjectsProjectTriggerProjectTriggerProjectTriggerResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectTriggerResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}