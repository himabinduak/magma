// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Csfb csfb configuration
//
// swagger:model csfb
type Csfb struct {

	// client
	Client *SctpClientConfigs `json:"client,omitempty"`
}

// Validate validates this csfb
func (m *Csfb) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClient(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Csfb) validateClient(formats strfmt.Registry) error {
	if swag.IsZero(m.Client) { // not required
		return nil
	}

	if m.Client != nil {
		if err := m.Client.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this csfb based on the context it is used
func (m *Csfb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateClient(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Csfb) contextValidateClient(ctx context.Context, formats strfmt.Registry) error {

	if m.Client != nil {
		if err := m.Client.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("client")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("client")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Csfb) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Csfb) UnmarshalBinary(b []byte) error {
	var res Csfb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}