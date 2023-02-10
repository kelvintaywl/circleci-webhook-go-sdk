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

// WebhookPayloadForResponse webhook payload for response
//
// swagger:model WebhookPayloadForResponse
type WebhookPayloadForResponse struct {
	WebhookBasePayload

	// Secret used to build an HMAC hash of the payload and passed as a header in the webhook request
	// Example: ****
	SigningSecret string `json:"signing_secret,omitempty"`

	// Whether to enforce TLS certificate verification when delivering the webhook
	// Example: true
	VerifyTLS bool `json:"verify_tls,omitempty"`
}

// UnmarshalJSON unmarshals this object from a JSON structure
func (m *WebhookPayloadForResponse) UnmarshalJSON(raw []byte) error {
	// AO0
	var aO0 WebhookBasePayload
	if err := swag.ReadJSON(raw, &aO0); err != nil {
		return err
	}
	m.WebhookBasePayload = aO0

	// AO1
	var dataAO1 struct {
		SigningSecret string `json:"signing_secret,omitempty"`

		VerifyTLS bool `json:"verify_tls,omitempty"`
	}
	if err := swag.ReadJSON(raw, &dataAO1); err != nil {
		return err
	}

	m.SigningSecret = dataAO1.SigningSecret

	m.VerifyTLS = dataAO1.VerifyTLS

	return nil
}

// MarshalJSON marshals this object to a JSON structure
func (m WebhookPayloadForResponse) MarshalJSON() ([]byte, error) {
	_parts := make([][]byte, 0, 2)

	aO0, err := swag.WriteJSON(m.WebhookBasePayload)
	if err != nil {
		return nil, err
	}
	_parts = append(_parts, aO0)
	var dataAO1 struct {
		SigningSecret string `json:"signing_secret,omitempty"`

		VerifyTLS bool `json:"verify_tls,omitempty"`
	}

	dataAO1.SigningSecret = m.SigningSecret

	dataAO1.VerifyTLS = m.VerifyTLS

	jsonDataAO1, errAO1 := swag.WriteJSON(dataAO1)
	if errAO1 != nil {
		return nil, errAO1
	}
	_parts = append(_parts, jsonDataAO1)
	return swag.ConcatJSON(_parts...), nil
}

// Validate validates this webhook payload for response
func (m *WebhookPayloadForResponse) Validate(formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WebhookBasePayload
	if err := m.WebhookBasePayload.Validate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validate this webhook payload for response based on the context it is used
func (m *WebhookPayloadForResponse) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	// validation for a type composition with WebhookBasePayload
	if err := m.WebhookBasePayload.ContextValidate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *WebhookPayloadForResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WebhookPayloadForResponse) UnmarshalBinary(b []byte) error {
	var res WebhookPayloadForResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}