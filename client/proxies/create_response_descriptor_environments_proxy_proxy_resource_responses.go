// Code generated by go-swagger; DO NOT EDIT.

package proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CreateResponseDescriptorEnvironmentsProxyProxyResourceReader is a Reader for the CreateResponseDescriptorEnvironmentsProxyProxyResource structure.
type CreateResponseDescriptorEnvironmentsProxyProxyResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewCreateResponseDescriptorEnvironmentsProxyProxyResourceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateResponseDescriptorEnvironmentsProxyProxyResourceCreated creates a CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated with default headers values
func NewCreateResponseDescriptorEnvironmentsProxyProxyResourceCreated() *CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated {
	return &CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated{}
}

/*CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated handles this case with default header values.

ProxyResource Created
*/
type CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated struct {
	Payload *models.ProxyResource
}

func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated) Error() string {
	return fmt.Sprintf("[POST /api/proxies][%d] createResponseDescriptorEnvironmentsProxyProxyResourceCreated  %+v", 201, o.Payload)
}

func (o *CreateResponseDescriptorEnvironmentsProxyProxyResourceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ProxyResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}