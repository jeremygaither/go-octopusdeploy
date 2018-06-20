// Code generated by go-swagger; DO NOT EDIT.

package project_groups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceReader is a Reader for the ModifyResponseDescriptorProjectsProjectGroupProjectGroupResource structure.
type ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK creates a ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK with default headers values
func NewModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK() *ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK {
	return &ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK{}
}

/*ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK handles this case with default header values.

ProjectGroupResource Modified
*/
type ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK struct {
	Payload *models.ProjectGroupResource
}

func (o *ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/projectgroups/{id}][%d] modifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK  %+v", 200, o.Payload)
}

func (o *ModifyResponseDescriptorProjectsProjectGroupProjectGroupResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProjectGroupResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}