// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ScannerRegistration Registration represents a named configuration for invoking a scanner via its adapter.
//
// swagger:model ScannerRegistration
type ScannerRegistration struct {

	// An optional value of the HTTP Authorization header sent with each request to the Scanner Adapter API.
	//
	// Example: Bearer: JWTTOKENGOESHERE
	AccessCredential string `json:"access_credential"`

	// Optional property to describe the name of the scanner registration
	// Example: Trivy
	Adapter string `json:"adapter,omitempty"`

	// Specify what authentication approach is adopted for the HTTP communications.
	// Supported types Basic", "Bearer" and api key header "X-ScannerAdapter-API-Key"
	//
	// Example: Bearer
	Auth string `json:"auth"`

	// The creation time of this registration
	// Format: date-time
	CreateTime strfmt.DateTime `json:"create_time,omitempty"`

	// An optional description of this registration.
	// Example: A free-to-use tool that scans container images for package vulnerabilities.\n
	Description string `json:"description"`

	// Indicate whether the registration is enabled or not
	Disabled *bool `json:"disabled"`

	// Indicate the healthy of the registration
	// Example: healthy
	Health string `json:"health,omitempty"`

	// Indicate if the registration is set as the system default one
	IsDefault *bool `json:"is_default"`

	// The name of this registration.
	// Example: Trivy
	Name string `json:"name,omitempty"`

	// Indicate if skip the certificate verification when sending HTTP requests
	SkipCertVerify *bool `json:"skip_certVerify"`

	// The update time of this registration
	// Format: date-time
	UpdateTime strfmt.DateTime `json:"update_time,omitempty"`

	// A base URL of the scanner adapter
	// Example: http://harbor-scanner-trivy:8080
	URL string `json:"url,omitempty"`

	// Indicate whether use internal registry addr for the scanner to pull content or not
	UseInternalAddr *bool `json:"use_internal_addr"`

	// The unique identifier of this registration.
	UUID string `json:"uuid,omitempty"`

	// Optional property to describe the vendor of the scanner registration
	// Example: CentOS
	Vendor string `json:"vendor,omitempty"`

	// Optional property to describe the version of the scanner registration
	// Example: 1.0.1
	Version string `json:"version,omitempty"`
}

// Validate validates this scanner registration
func (m *ScannerRegistration) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdateTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ScannerRegistration) validateCreateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.CreateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("create_time", "body", "date-time", m.CreateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ScannerRegistration) validateUpdateTime(formats strfmt.Registry) error {
	if swag.IsZero(m.UpdateTime) { // not required
		return nil
	}

	if err := validate.FormatOf("update_time", "body", "date-time", m.UpdateTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this scanner registration based on context it is used
func (m *ScannerRegistration) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ScannerRegistration) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ScannerRegistration) UnmarshalBinary(b []byte) error {
	var res ScannerRegistration
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
