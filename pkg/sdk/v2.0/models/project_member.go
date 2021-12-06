// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ProjectMember project member
//
// swagger:model ProjectMember
type ProjectMember struct {

	// member group
	MemberGroup *UserGroup `json:"member_group,omitempty"`

	// member user
	MemberUser *UserEntity `json:"member_user,omitempty"`

	// The role id 1 for projectAdmin, 2 for developer, 3 for guest, 4 for maintainer
	RoleID int64 `json:"role_id,omitempty"`
}

// Validate validates this project member
func (m *ProjectMember) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateMemberGroup(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMemberUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectMember) validateMemberGroup(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberGroup) { // not required
		return nil
	}

	if m.MemberGroup != nil {
		if err := m.MemberGroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member_group")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectMember) validateMemberUser(formats strfmt.Registry) error {

	if swag.IsZero(m.MemberUser) { // not required
		return nil
	}

	if m.MemberUser != nil {
		if err := m.MemberUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("member_user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectMember) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectMember) UnmarshalBinary(b []byte) error {
	var res ProjectMember
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
