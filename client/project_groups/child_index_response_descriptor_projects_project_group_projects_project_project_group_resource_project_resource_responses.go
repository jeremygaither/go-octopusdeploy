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

// ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceReader is a Reader for the ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResource structure.
type ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK creates a ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK with default headers values
func NewChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK() *ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK {
	return &ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK{}
}

/*ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK handles this case with default header values.

Load ProjectGroupResource collection
*/
type ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK struct {
	Payload *models.ResourceCollectionProjectGroupResource
}

func (o *ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/projectgroups/{id}/projects][%d] childIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK  %+v", 200, o.Payload)
}

func (o *ChildIndexResponseDescriptorProjectsProjectGroupProjectsProjectProjectGroupResourceProjectResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionProjectGroupResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}