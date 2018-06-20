// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EndpointResource endpoint resource
// swagger:model EndpointResource
type EndpointResource struct {

	// communication style
	// Read Only: true
	// Enum: [None TentaclePassive TentacleActive Ssh OfflineDrop AzureWebApp Ftp AzureCloudService]
	CommunicationStyle int32 `json:"CommunicationStyle,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`
}

// Validate validates this endpoint resource
func (m *EndpointResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCommunicationStyle(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastModifiedOn(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var endpointResourceTypeCommunicationStylePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`["None","TentaclePassive","TentacleActive","Ssh","OfflineDrop","AzureWebApp","Ftp","AzureCloudService"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		endpointResourceTypeCommunicationStylePropEnum = append(endpointResourceTypeCommunicationStylePropEnum, v)
	}
}

// prop value enum
func (m *EndpointResource) validateCommunicationStyleEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, endpointResourceTypeCommunicationStylePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EndpointResource) validateCommunicationStyle(formats strfmt.Registry) error {

	if swag.IsZero(m.CommunicationStyle) { // not required
		return nil
	}

	// value enum
	if err := m.validateCommunicationStyleEnum("CommunicationStyle", "body", m.CommunicationStyle); err != nil {
		return err
	}

	return nil
}

func (m *EndpointResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EndpointResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EndpointResource) UnmarshalBinary(b []byte) error {
	var res EndpointResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}