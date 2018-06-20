// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// ReleaseTemplatePackage release template package
// swagger:model ReleaseTemplatePackage
type ReleaseTemplatePackage struct {

	// action name
	ActionName string `json:"ActionName,omitempty"`

	// feed Id
	FeedID string `json:"FeedId,omitempty"`

	// feed name
	FeedName string `json:"FeedName,omitempty"`

	// is resolvable
	IsResolvable bool `json:"IsResolvable,omitempty"`

	// nu get feed Id
	NuGetFeedID string `json:"NuGetFeedId,omitempty"`

	// nu get feed name
	NuGetFeedName string `json:"NuGetFeedName,omitempty"`

	// nu get package Id
	NuGetPackageID string `json:"NuGetPackageId,omitempty"`

	// package Id
	PackageID string `json:"PackageId,omitempty"`

	// project name
	ProjectName string `json:"ProjectName,omitempty"`

	// step name
	StepName string `json:"StepName,omitempty"`

	// version selected last release
	VersionSelectedLastRelease string `json:"VersionSelectedLastRelease,omitempty"`
}

// Validate validates this release template package
func (m *ReleaseTemplatePackage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ReleaseTemplatePackage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReleaseTemplatePackage) UnmarshalBinary(b []byte) error {
	var res ReleaseTemplatePackage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}