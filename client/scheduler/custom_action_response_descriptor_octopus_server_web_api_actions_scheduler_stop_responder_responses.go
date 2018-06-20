// Code generated by go-swagger; DO NOT EDIT.

package scheduler

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK handles this case with default header values.

OK
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/scheduler/stop][%d] customActionResponseDescriptorOctopusServerWebApiActionsSchedulerStopResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsSchedulerStopResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}