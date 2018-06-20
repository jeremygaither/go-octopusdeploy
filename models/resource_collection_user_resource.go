// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ResourceCollectionUserResource resource collection user resource
// swagger:model ResourceCollection[UserResource]
type ResourceCollectionUserResource struct {

	// Id
	ID string `json:"Id,omitempty"`

	// item type
	// Read Only: true
	ItemType string `json:"ItemType,omitempty"`

	// items
	Items []*UserResource `json:"Items"`

	// items per page
	ItemsPerPage int32 `json:"ItemsPerPage,omitempty"`

	// last modified by
	LastModifiedBy string `json:"LastModifiedBy,omitempty"`

	// last modified on
	// Format: date-time
	LastModifiedOn strfmt.DateTime `json:"LastModifiedOn,omitempty"`

	// last page number
	// Read Only: true
	LastPageNumber int32 `json:"LastPageNumber,omitempty"`

	// links
	Links map[string]string `json:"Links,omitempty"`

	// number of pages
	// Read Only: true
	NumberOfPages int32 `json:"NumberOfPages,omitempty"`

	// total results
	TotalResults int32 `json:"TotalResults,omitempty"`
}

// Validate validates this resource collection user resource
func (m *ResourceCollectionUserResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateItems(formats); err != nil {
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

func (m *ResourceCollectionUserResource) validateItems(formats strfmt.Registry) error {

	if swag.IsZero(m.Items) { // not required
		return nil
	}

	for i := 0; i < len(m.Items); i++ {
		if swag.IsZero(m.Items[i]) { // not required
			continue
		}

		if m.Items[i] != nil {
			if err := m.Items[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Items" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *ResourceCollectionUserResource) validateLastModifiedOn(formats strfmt.Registry) error {

	if swag.IsZero(m.LastModifiedOn) { // not required
		return nil
	}

	if err := validate.FormatOf("LastModifiedOn", "body", "date-time", m.LastModifiedOn.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ResourceCollectionUserResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ResourceCollectionUserResource) UnmarshalBinary(b []byte) error {
	var res ResourceCollectionUserResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}