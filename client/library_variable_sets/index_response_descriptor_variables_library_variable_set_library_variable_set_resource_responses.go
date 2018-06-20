// Code generated by go-swagger; DO NOT EDIT.

package library_variable_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceReader is a Reader for the IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResource structure.
type IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK creates a IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK with default headers values
func NewIndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK() *IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK {
	return &IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK{}
}

/*IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK handles this case with default header values.

Load LibraryVariableSetResource collection
*/
type IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK struct {
	Payload *models.ResourceCollectionLibraryVariableSetResource
}

func (o *IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/libraryvariablesets][%d] indexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK  %+v", 200, o.Payload)
}

func (o *IndexResponseDescriptorVariablesLibraryVariableSetLibraryVariableSetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ResourceCollectionLibraryVariableSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}