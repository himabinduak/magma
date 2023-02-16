// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// FegNetworkID Name of the federated network serving this LTE network. Blank for non federated
// Example: feg_id_only_for_federated_network
//
// swagger:model fegNetworkId
type FegNetworkID string

// Validate validates this feg network Id
func (m FegNetworkID) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this feg network Id based on context it is used
func (m FegNetworkID) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}