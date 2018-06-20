// Code generated by go-swagger; DO NOT EDIT.

package status

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponder structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK handles this case with default header values.

OK - Default
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK) Error() string {
	return fmt.Sprintf("[GET /api/serverstatus/system-report][%d] customActionResponseDescriptorOctopusServerWebApiActionsStatusSystemReportResponderOK ", 200)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsStatusSystemReportResponderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}