// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// IdentityResource identity resource
// swagger:model IdentityResource
type IdentityResource struct {

	// claims
	Claims IdentityResourceClaims `json:"Claims,omitempty"`

	// identity provider name
	IdentityProviderName string `json:"IdentityProviderName,omitempty"`
}

// Validate validates this identity resource
func (m *IdentityResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClaims(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IdentityResource) validateClaims(formats strfmt.Registry) error {

	if swag.IsZero(m.Claims) { // not required
		return nil
	}

	if err := m.Claims.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("Claims")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IdentityResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IdentityResource) UnmarshalBinary(b []byte) error {
	var res IdentityResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}