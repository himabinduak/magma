// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SuciProfile SUCI profile for a network. These attributes are defined as per the SUCI encryption/decryption mechanisms specified in the 3GPP TS 33.501.
//
// swagger:model suciProfile
type SuciProfile struct {

	// HPLMN Private Key
	// Format: byte
	HomeNetworkPrivateKey strfmt.Base64 `json:"home_network_private_key,omitempty"`

	// HPLMN Public Key
	// Format: byte
	HomeNetworkPublicKey strfmt.Base64 `json:"home_network_public_key,omitempty"`

	// HPLMN Public Key Identifier
	// Maximum: 255
	// Minimum: 0
	HomeNetworkPublicKeyIdentifier uint32 `json:"home_network_public_key_identifier,omitempty"`

	// ECIES Protection scheme
	// Enum: [ProfileA ProfileB]
	ProtectionScheme string `json:"protection_scheme,omitempty"`
}

// Validate validates this suci profile
func (m *SuciProfile) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHomeNetworkPublicKeyIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtectionScheme(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SuciProfile) validateHomeNetworkPublicKeyIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.HomeNetworkPublicKeyIdentifier) { // not required
		return nil
	}

	if err := validate.MinimumUint("home_network_public_key_identifier", "body", uint64(m.HomeNetworkPublicKeyIdentifier), 0, false); err != nil {
		return err
	}

	if err := validate.MaximumUint("home_network_public_key_identifier", "body", uint64(m.HomeNetworkPublicKeyIdentifier), 255, false); err != nil {
		return err
	}

	return nil
}

var suciProfileTypeProtectionSchemePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ProfileA","ProfileB"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		suciProfileTypeProtectionSchemePropEnum = append(suciProfileTypeProtectionSchemePropEnum, v)
	}
}

const (

	// SuciProfileProtectionSchemeProfileA captures enum value "ProfileA"
	SuciProfileProtectionSchemeProfileA string = "ProfileA"

	// SuciProfileProtectionSchemeProfileB captures enum value "ProfileB"
	SuciProfileProtectionSchemeProfileB string = "ProfileB"
)

// prop value enum
func (m *SuciProfile) validateProtectionSchemeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, suciProfileTypeProtectionSchemePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SuciProfile) validateProtectionScheme(formats strfmt.Registry) error {
	if swag.IsZero(m.ProtectionScheme) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtectionSchemeEnum("protection_scheme", "body", m.ProtectionScheme); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this suci profile based on context it is used
func (m *SuciProfile) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *SuciProfile) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SuciProfile) UnmarshalBinary(b []byte) error {
	var res SuciProfile
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}