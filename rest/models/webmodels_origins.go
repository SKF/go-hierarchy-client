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

// WebmodelsOrigins webmodels origins
//
// swagger:model webmodels.Origins
type WebmodelsOrigins struct {

	// origins
	Origins []Origin `json:"origins"`
}

// Validate validates this webmodels origins
func (m *WebmodelsOrigins) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateOrigins(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WebmodelsOrigins) validateOrigins(formats strfmt.Registry) error {

	if swag.IsZero(m.Origins) { // not required
		return nil
	}

	for i := 0; i < len(m.Origins); i++ {

		if err := m.Origins[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origins" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *WebmodelsOrigins) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebmodelsOrigins) UnmarshalBinary(b []byte) error {
	var res WebmodelsOrigins
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
