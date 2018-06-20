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

// ModifyResponseDescriptorProjectsChannelChannelResourceReader is a Reader for the ModifyResponseDescriptorProjectsChannelChannelResource structure.
type ModifyResponseDescriptorProjectsChannelChannelResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ModifyResponseDescriptorProjectsChannelChannelResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewModifyResponseDescriptorProjectsChannelChannelResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewModifyResponseDescriptorProjectsChannelChannelResourceOK creates a ModifyResponseDescriptorProjectsChannelChannelResourceOK with default headers values
func NewModifyResponseDescriptorProjectsChannelChannelResourceOK() *ModifyResponseDescriptorProjectsChannelChannelResourceOK {
	return &ModifyResponseDescriptorProjectsChannelChannelResourceOK{}
}

/*ModifyResponseDescriptorProjectsChannelChannelResourceOK handles this case with default header values.

ChannelResource Modified
*/
type ModifyResponseDescriptorProjectsChannelChannelResourceOK struct {
	Payload *models.ChannelResource
}

func (o *ModifyResponseDescriptorProjectsChannelChannelResourceOK) Error() string {
	return fmt.Sprintf("[PUT /api/channels/{id}][%d] modifyResponseDescriptorProjectsChannelChannelResourceOK  %+v", 200, o.Payload)
}

func (o *ModifyResponseDescriptorProjectsChannelChannelResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ChannelResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}