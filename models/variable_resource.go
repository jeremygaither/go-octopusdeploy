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

// VariableResource variable resource
// swagger:model VariableResource
type VariableResource struct {

	// description
	Description string `json:"Description,omitempty"`

	// Id
	ID string `json:"Id,omitempty"`

	// is editable
	IsEditable bool `json:"IsEditable,omitempty"`

	// is sensitive
	IsSensitive bool `json:"IsSensitive,omitempty"`

	// name
	Name string `json:"Name,omitempty"`

	// prompt
	Prompt *VariablePromptOptions `json:"Prompt,omitempty"`

	// scope
	Scope *VariableResourceScope `json:"Scope,omitempty"`

	// type
	// Enum: [String Sensitive Certificate AmazonWebServicesAccount]
	Type int32 `json:"Type,omitempty"`

	// value
	Value string `json:"Value,omitempty"`
}

// Validate validates this variable resource
func (m *VariableResource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePrompt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *VariableResource) validatePrompt(formats strfmt.Registry) error {

	if swag.IsZero(m.Prompt) { // not required
		return nil
	}

	if m.Prompt != nil {
		if err := m.Prompt.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Prompt")
			}
			return err
		}
	}

	return nil
}

func (m *VariableResource) validateScope(formats strfmt.Registry) error {

	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	if m.Scope != nil {
		if err := m.Scope.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Scope")
			}
			return err
		}
	}

	return nil
}

var variableResourceTypeTypePropEnum []interface{}

func init() {
	var res []int32
	if err := json.Unmarshal([]byte(`["String","Sensitive","Certificate","AmazonWebServicesAccount"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		variableResourceTypeTypePropEnum = append(variableResourceTypeTypePropEnum, v)
	}
}

// prop value enum
func (m *VariableResource) validateTypeEnum(path, location string, value int32) error {
	if err := validate.Enum(path, location, value, variableResourceTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *VariableResource) validateType(formats strfmt.Registry) error {

	if swag.IsZero(m.Type) { // not required
		return nil
	}

	// value enum
	if err := m.validateTypeEnum("Type", "body", m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *VariableResource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *VariableResource) UnmarshalBinary(b []byte) error {
	var res VariableResource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}