// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorTagSetsTagSetTagSetResourceReader is a Reader for the IndexResponseDescriptorTagSetsTagSetTagSetResource structure.
type IndexResponseDescriptorTagSetsTagSetTagSetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorTagSetsTagSetTagSetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorTagSetsTagSetTagSetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorTagSetsTagSetTagSetResourceOK creates a IndexResponseDescriptorTagSetsTagSetTagSetResourceOK with default headers values
func NewIndexResponseDescriptorTagSetsTagSetTagSetResourceOK() *IndexResponseDescriptorTagSetsTagSetTagSetResourceOK {
	return &IndexResponseDescriptorTagSetsTagSetTagSetResourceOK{}
}

/*IndexResponseDescriptorTagSetsTagSetTagSetResourceOK handles this case with default header values.

Load TagSetResource collection
*/
type IndexResponseDescriptorTagSetsTagSetTagSetResourceOK struct {
	Payload *models.ResourceCollectionTagSetResource
}

func (o *IndexResponseDescriptorTagSetsTagSetTagSetResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/tagsets][%d] indexResponseDescriptorTagSetsTagSetTagSetResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorTagSetsTagSetTagSetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionTagSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}