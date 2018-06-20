// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// DocumentTypeDocument document type document
// swagger:model DocumentTypeDocument
type DocumentTypeDocument struct {

	// Id
	ID string `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this document type document
func (m *DocumentTypeDocument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DocumentTypeDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DocumentTypeDocument) UnmarshalBinary(b []byte) error {
	var res DocumentTypeDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}