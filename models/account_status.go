// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// AccountStatus Status of the Account.
// "Ready": The account is clean and ready for lease
// "NotReady": The account is in "dirty" state, and needs to be reset before it may be leased.
// "Leased": The account is leased to a principal
//
// swagger:model accountStatus
type AccountStatus string

const (

	// AccountStatusReady captures enum value "Ready"
	AccountStatusReady AccountStatus = "Ready"

	// AccountStatusNotReady captures enum value "NotReady"
	AccountStatusNotReady AccountStatus = "NotReady"

	// AccountStatusLeased captures enum value "Leased"
	AccountStatusLeased AccountStatus = "Leased"

	// AccountStatusOrphaned captures enum value "Orphaned"
	AccountStatusOrphaned AccountStatus = "Orphaned"
)

// for schema
var accountStatusEnum []interface{}

func init() {
	var res []AccountStatus
	if err := json.Unmarshal([]byte(`["Ready","NotReady","Leased","Orphaned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		accountStatusEnum = append(accountStatusEnum, v)
	}
}

func (m AccountStatus) validateAccountStatusEnum(path, location string, value AccountStatus) error {
	if err := validate.Enum(path, location, value, accountStatusEnum); err != nil {
		return err
	}
	return nil
}

// Validate validates this account status
func (m AccountStatus) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateAccountStatusEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
