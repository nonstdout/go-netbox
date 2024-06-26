// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+nonstdout@users.noreply.github.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NestedModule nested module
//
// swagger:model NestedModule
type NestedModule struct {

	// device
	Device *NestedDevice `json:"device,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// module bay
	ModuleBay *ModuleNestedModuleBay `json:"module_bay,omitempty"`

	// module type
	ModuleType *NestedModuleType `json:"module_type,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this nested module
func (m *NestedModule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModuleBay(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateModuleType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedModule) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(m.Device) { // not required
		return nil
	}

	if m.Device != nil {
		if err := m.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) validateModuleBay(formats strfmt.Registry) error {
	if swag.IsZero(m.ModuleBay) { // not required
		return nil
	}

	if m.ModuleBay != nil {
		if err := m.ModuleBay.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_bay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_bay")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) validateModuleType(formats strfmt.Registry) error {
	if swag.IsZero(m.ModuleType) { // not required
		return nil
	}

	if m.ModuleType != nil {
		if err := m.ModuleType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_type")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this nested module based on the context it is used
func (m *NestedModule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModuleBay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateModuleType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *NestedModule) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if m.Device != nil {
		if err := m.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *NestedModule) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *NestedModule) contextValidateModuleBay(ctx context.Context, formats strfmt.Registry) error {

	if m.ModuleBay != nil {
		if err := m.ModuleBay.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_bay")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_bay")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) contextValidateModuleType(ctx context.Context, formats strfmt.Registry) error {

	if m.ModuleType != nil {
		if err := m.ModuleType.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("module_type")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("module_type")
			}
			return err
		}
	}

	return nil
}

func (m *NestedModule) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NestedModule) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NestedModule) UnmarshalBinary(b []byte) error {
	var res NestedModule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
