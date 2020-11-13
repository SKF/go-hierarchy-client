// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AssetType asset type
//
// swagger:model AssetType
type AssetType struct {

	// code
	Code string `json:"code,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// sequences
	Sequences []*AssetSequence `json:"sequences"`
}

// Validate validates this asset type
func (m *AssetType) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSequences(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssetType) validateSequences(formats strfmt.Registry) error {

	if swag.IsZero(m.Sequences) { // not required
		return nil
	}

	for i := 0; i < len(m.Sequences); i++ {
		if swag.IsZero(m.Sequences[i]) { // not required
			continue
		}

		if m.Sequences[i] != nil {
			if err := m.Sequences[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("sequences" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssetType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssetType) UnmarshalBinary(b []byte) error {
	var res AssetType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}