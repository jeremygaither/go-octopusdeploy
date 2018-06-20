// Code generated by go-swagger; DO NOT EDIT.

package tag_sets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK handles this case with default header values.

OK - Default
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK) Error() string {
	return fmt.Sprintf("[PUT /api/tagsets/sortorder][%d] customActionResponseDescriptorOctopusServerWebApiActionsSortTagSetsResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSortTagSetsResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}