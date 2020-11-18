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

// WebmodelsAssetClasses webmodels asset classes
//
// swagger:model webmodels.AssetClasses
type WebmodelsAssetClasses struct {

	// classes
	Classes []*AssetClass `json:"classes"`
}

// Validate validates this webmodels asset classes
func (m *WebmodelsAssetClasses) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClasses(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebmodelsAssetClasses) validateClasses(formats strfmt.Registry) error {

	if swag.IsZero(m.Classes) { // not required
		return nil
	}

	for i := 0; i < len(m.Classes); i++ {
		if swag.IsZero(m.Classes[i]) { // not required
			continue
		}

		if m.Classes[i] != nil {
			if err := m.Classes[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("classes" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebmodelsAssetClasses) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebmodelsAssetClasses) UnmarshalBinary(b []byte) error {
	var res WebmodelsAssetClasses
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
