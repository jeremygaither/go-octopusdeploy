// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// CommunityActionTemplateResource community action template resource
// swagger:model CommunityActionTemplateResource
type CommunityActionTemplateResource struct {

	// author
	Author string `json:"Author,omitempty"`

	// description
	Description string `json:"Description,omitempty"`

	// history Url
	HistoryURL string `json:"HistoryUrl,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// parameters
	// Read Only: true
	Parameters []*ActionTemplateParameterResource `json:"Parameters"`

	// properties
	Properties CommunityActionTemplateResourceProperties `json:"Properties,omitempty"`

	// type
	Type string `json:"Type,omitempty"`

	// version
	Version int32 `json:"Version,omitempty"`

	// website
	Website string `json:"Website,omitempty"`
}

// Validate validates this community action template resource
func (m *CommunityActionTemplateResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProperties(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CommunityActionTemplateResource) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *CommunityActionTemplateResource) validateProperties(formats strfmt.Registry) error {

	if swag.IsZero(m.Properties) { // not required
		return nil
	}

	if err := m.Properties.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Properties")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CommunityActionTemplateResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CommunityActionTemplateResource) UnmarshalBinary(b []byte) error {
	var res CommunityActionTemplateResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}