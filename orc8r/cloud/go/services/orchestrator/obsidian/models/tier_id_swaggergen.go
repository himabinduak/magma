// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// TierID tier id
// Example: default
//
// swagger:model tier_id
type TierID string

// Validate validates this tier id
func (m TierID) Validate(formats strfmt.Registry) error {
	var res []error

	if err := validate.Pattern("", "body", string(m), `^[a-z][\da-z_]+$`); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this tier id based on context it is used
func (m TierID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}