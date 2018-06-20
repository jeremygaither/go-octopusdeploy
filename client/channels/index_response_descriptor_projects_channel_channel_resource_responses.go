// Code generated by go-swagger; DO NOT EDIT.

package channels

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorProjectsChannelChannelResourceReader is a Reader for the IndexResponseDescriptorProjectsChannelChannelResource structure.
type IndexResponseDescriptorProjectsChannelChannelResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorProjectsChannelChannelResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorProjectsChannelChannelResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorProjectsChannelChannelResourceOK creates a IndexResponseDescriptorProjectsChannelChannelResourceOK with default headers values
func NewIndexResponseDescriptorProjectsChannelChannelResourceOK() *IndexResponseDescriptorProjectsChannelChannelResourceOK {
	return &IndexResponseDescriptorProjectsChannelChannelResourceOK{}
}

/*IndexResponseDescriptorProjectsChannelChannelResourceOK handles this case with default header values.

Load ChannelResource collection
*/
type IndexResponseDescriptorProjectsChannelChannelResourceOK struct {
	Payload *models.ResourceCollectionChannelResource
}

func (o *IndexResponseDescriptorProjectsChannelChannelResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/channels][%d] indexResponseDescriptorProjectsChannelChannelResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorProjectsChannelChannelResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionChannelResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}