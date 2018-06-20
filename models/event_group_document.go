// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EventGroupDocument event group document
// swagger:model EventGroupDocument
type EventGroupDocument struct {

	// event categories
	EventCategories []string `json:"EventCategories"`

	// Id
	ID string `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this event group document
func (m *EventGroupDocument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventGroupDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventGroupDocument) UnmarshalBinary(b []byte) error {
	var res EventGroupDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}