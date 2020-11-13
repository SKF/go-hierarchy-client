// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Origin origin
//
// swagger:model Origin
type Origin struct {

	// Origin identity
	ID string `json:"id,omitempty"`

	// Origin provider
	Provider string `json:"provider,omitempty"`

	// Origin type
	Type string `json:"type,omitempty"`
}

// Validate validates this origin
func (m *Origin) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Origin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Origin) UnmarshalBinary(b []byte) error {
	var res Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}