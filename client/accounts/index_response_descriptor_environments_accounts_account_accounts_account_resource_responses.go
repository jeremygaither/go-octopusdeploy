// Code generated by go-swagger; DO NOT EDIT.

package accounts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader is a Reader for the IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResource structure.
type IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK creates a IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK with default headers values
func NewIndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK() *IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK {
	return &IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK{}
}

/*IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK handles this case with default header values.

Load AccountResource collection
*/
type IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK struct {
	Payload *models.ResourceCollectionAccountResource
}

func (o *IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/accounts][%d] indexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorEnvironmentsAccountsAccountAccountsAccountResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionAccountResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}