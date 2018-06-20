// Code generated by go-swagger; DO NOT EDIT.

package releases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceReader is a Reader for the ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResource structure.
type ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK creates a ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK with default headers values
func NewChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK() *ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK {
	return &ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK{}
}

/*ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK handles this case with default header values.

Load ProjectResource collection
*/
type ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK struct {
	Payload *models.ResourceCollectionProjectResource
}

func (o *ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/projects/{id}/releases][%d] childIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK  %+v", 200, o.Payload)
}

func (o *ChildIndexResponseDescriptorProjectsProjectProjectsReleaseProjectResourceReleaseResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionProjectResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}