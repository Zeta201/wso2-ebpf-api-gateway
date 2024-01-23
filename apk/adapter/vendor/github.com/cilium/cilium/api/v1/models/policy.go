// Code generated by go-swagger; DO NOT EDIT.

// Copyright Authors of Cilium
// SPDX-License-Identifier: Apache-2.0

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Policy Policy definition
//
// swagger:model Policy
type Policy struct {

	// Policy definition as JSON.
	Policy string `json:"policy,omitempty"`

	// Revision number of the policy. Incremented each time the policy is
	// changed in the agent's repository
	//
	Revision int64 `json:"revision,omitempty"`
}

// Validate validates this policy
func (m *Policy) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this policy based on context it is used
func (m *Policy) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Policy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Policy) UnmarshalBinary(b []byte) error {
	var res Policy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
