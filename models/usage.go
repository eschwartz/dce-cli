// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Usage usage cost of the aws account from start date to end date
// swagger:model usage
type Usage struct {

	// accountId of the AWS account
	AccountID string `json:"accountId,omitempty"`

	// usage cost Amount of AWS account for given period
	CostAmount float64 `json:"costAmount,omitempty"`

	// usage cost currency
	CostCurrency string `json:"costCurrency,omitempty"`

	// usage end date as Epoch Timestamp
	EndDate float64 `json:"endDate,omitempty"`

	// principalId of the user who owns the lease of the AWS account
	//
	PrincipalID string `json:"principalId,omitempty"`

	// usage start date as Epoch Timestamp
	StartDate float64 `json:"startDate,omitempty"`

	// ttl attribute as Epoch Timestamp
	TimeToLive float64 `json:"timeToLive,omitempty"`
}

// Validate validates this usage
func (m *Usage) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Usage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Usage) UnmarshalBinary(b []byte) error {
	var res Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
