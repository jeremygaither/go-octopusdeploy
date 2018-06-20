// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InvitationResource invitation resource
// swagger:model InvitationResource
type InvitationResource struct {

	// add to team ids
	// Required: true
	AddToTeamIds []string `json:"AddToTeamIds"`

	// expires
	// Format: date-time
	Expires strfmt.DateTime `json:"Expires,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// invitation code
	InvitationCode string `json:"InvitationCode,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`
}

// Validate validates this invitation resource
func (m *InvitationResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddToTeamIds(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpires(formats); err != nil {
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

func (m *InvitationResource) validateAddToTeamIds(formats strfmt.Registry) error {

	if err := validate.Required("AddToTeamIds", "body", m.AddToTeamIds); err != nil {
		return err
	}

	return nil
}

func (m *InvitationResource) validateExpires(formats strfmt.Registry) error {

	if swag.IsZero(m.Expires) { // not required
		return nil
	}

	if err := validate.FormatOf("Expires", "body", "date-time", m.Expires.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvitationResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InvitationResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvitationResource) UnmarshalBinary(b []byte) error {
	var res InvitationResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}