// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// EventCategoryDocument event category document
// swagger:model EventCategoryDocument
type EventCategoryDocument struct {

	// Id
	ID string `json:"Id,omitempty"`

	// name
	Name string `json:"Name,omitempty"`
}

// Validate validates this event category document
func (m *EventCategoryDocument) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EventCategoryDocument) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EventCategoryDocument) UnmarshalBinary(b []byte) error {
	var res EventCategoryDocument
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}