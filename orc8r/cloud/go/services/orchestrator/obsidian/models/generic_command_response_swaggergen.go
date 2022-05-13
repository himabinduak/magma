// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GenericCommandResponse generic command response
//
// swagger:model generic_command_response
type GenericCommandResponse struct {

	// response
	// Example: {}
	Response map[string]interface{} `json:"response,omitempty"`
}

// Validate validates this generic command response
func (m *GenericCommandResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this generic command response based on context it is used
func (m *GenericCommandResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GenericCommandResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GenericCommandResponse) UnmarshalBinary(b []byte) error {
	var res GenericCommandResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}