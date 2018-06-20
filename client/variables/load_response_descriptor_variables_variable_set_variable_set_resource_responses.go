// Code generated by go-swagger; DO NOT EDIT.

package variables

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// LoadResponseDescriptorVariablesVariableSetVariableSetResourceReader is a Reader for the LoadResponseDescriptorVariablesVariableSetVariableSetResource structure.
type LoadResponseDescriptorVariablesVariableSetVariableSetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LoadResponseDescriptorVariablesVariableSetVariableSetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewLoadResponseDescriptorVariablesVariableSetVariableSetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewLoadResponseDescriptorVariablesVariableSetVariableSetResourceOK creates a LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK with default headers values
func NewLoadResponseDescriptorVariablesVariableSetVariableSetResourceOK() *LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK {
	return &LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK{}
}

/*LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK handles this case with default header values.

Load VariableSetResource by ID
*/
type LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK struct {
	Payload *models.VariableSetResource
}

func (o *LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK) Error() string {
	return fmt.Sprintf("[GET /api/variables/{id}][%d] loadResponseDescriptorVariablesVariableSetVariableSetResourceOK  %+v", 200, o.Payload)
}

func (o *LoadResponseDescriptorVariablesVariableSetVariableSetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VariableSetResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}