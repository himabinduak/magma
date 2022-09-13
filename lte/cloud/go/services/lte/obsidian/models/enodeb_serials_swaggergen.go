// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// EnodebSerials enodeb serials
// Example: ["SN1234567890","SN09876554321"]
//
// swagger:model enodeb_serials
type EnodebSerials []string

// Validate validates this enodeb serials
func (m EnodebSerials) Validate(formats strfmt.Registry) error {
	var res []error

	for i := 0; i < len(m); i++ {

		if err := validate.MinLength(strconv.Itoa(i), "body", m[i], 1); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this enodeb serials based on context it is used
func (m EnodebSerials) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}