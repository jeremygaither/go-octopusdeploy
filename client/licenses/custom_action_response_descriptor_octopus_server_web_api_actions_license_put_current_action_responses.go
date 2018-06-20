// Code generated by go-swagger; DO NOT EDIT.

package licenses

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/MattHodge/go-octopusdeploy/models"
)

// CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionReader is a Reader for the CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentAction structure.
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK() *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK handles this case with default header values.

LicenseResource resource returned
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK struct {
	Payload *models.LicenseResource
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK) Error() string {
	return fmt.Sprintf("[PUT /api/licenses/licenses-current][%d] customActionResponseDescriptorOctopusServerWebApiActionsLicensePutCurrentActionOK  %+v", 200, o.Payload)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LicenseResource)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest creates a CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest with default headers values
func NewCustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest() *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest {
	return &CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest{}
}

/*CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest handles this case with default header values.

Please provide the text of the license.
The license is not compatible with this version of Octopus deploy.
The license could not be updated because the server is not compliant, so using the license would disable functionality.
The license you've provided doesn't seem to be valid XML. You should paste in the entire XML document, starting with <License> and ending at </License>.
No request body was supplied.
*/
type CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest struct {
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /api/licenses/licenses-current][%d] customActionResponseDescriptorOctopusServerWebApiActionsLicensePutCurrentActionBadRequest ", 400)
}

func (o *CustomActionResponseDescriptorOctopusServerWebAPIActionsLicensePutCurrentActionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}