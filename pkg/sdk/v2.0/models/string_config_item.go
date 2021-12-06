// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// StringConfigItem string config item
//
// swagger:model StringConfigItem
type StringConfigItem struct {

	// The configure item can be updated or not
	Editable bool `json:"editable"`

	// The string value of current config item
	Value string `json:"value"`
}

// Validate validates this string config item
func (m *StringConfigItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *StringConfigItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StringConfigItem) UnmarshalBinary(b []byte) error {
	var res StringConfigItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
