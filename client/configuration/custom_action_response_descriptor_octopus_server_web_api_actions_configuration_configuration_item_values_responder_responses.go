// Code generated by go-swagger; DO NOT EDIT.

package configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK handles this case with default header values.

Object resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK struct {
	Payload CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOKBody
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/configuration/{id}/values][%d] customActionResponseDescriptorOctopusServerWebApiActionsConfigurationConfigurationItemValuesResponderOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest handles this case with default header values.

No id parameter was provided.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest) Error() string {
	return fmt.Sprintf("[GET /api/configuration/{id}/values][%d] customActionResponseDescriptorOctopusServerWebApiActionsConfigurationConfigurationItemValuesResponderBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOKBody custom action response descriptor octopus server web API actions configuration configuration item values responder o k body
swagger:model CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOKBody
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsConfigurationConfigurationItemValuesResponderOKBody interface{}