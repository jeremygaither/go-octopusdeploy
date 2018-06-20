// Code generated by go-swagger; DO NOT EDIT.

package subscription

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceReader is a Reader for the IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResource structure.
type IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK creates a IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK with default headers values
func NewIndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK() *IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK {
	return &IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK{}
}

/*IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK handles this case with default header values.

Load SubscriptionResource collection
*/
type IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK struct {
	Payload *models.ResourceCollectionSubscriptionResource
}

func (o *IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/subscriptions][%d] indexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorSubscriptionsSubscriptionSubscriptionResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionSubscriptionResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}