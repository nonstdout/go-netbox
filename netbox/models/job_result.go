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
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// JobResult job result
//
// swagger:model JobResult
type JobResult struct {

	// Completed
	// Format: date-time
	Completed *strfmt.DateTime `json:"completed,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created strfmt.DateTime `json:"created,omitempty"`

	// Data
	Data interface{} `json:"data,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Interval
	//
	// Recurrence interval (in minutes)
	// Maximum: 2.147483647e+09
	// Minimum: 1
	Interval *int64 `json:"interval,omitempty"`

	// Job id
	// Required: true
	// Format: uuid
	JobID *strfmt.UUID `json:"job_id"`

	// Name
	// Required: true
	// Max Length: 255
	// Min Length: 1
	Name *string `json:"name"`

	// Obj type
	// Read Only: true
	ObjType string `json:"obj_type,omitempty"`

	// Scheduled
	// Format: date-time
	Scheduled *strfmt.DateTime `json:"scheduled,omitempty"`

	// Started
	// Format: date-time
	Started *strfmt.DateTime `json:"started,omitempty"`

	// status
	Status *JobResultStatus `json:"status,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// user
	User *NestedUser `json:"user,omitempty"`
}

// Validate validates this job result
func (m *JobResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCompleted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterval(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScheduled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStarted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobResult) validateCompleted(formats strfmt.Registry) error {
	if swag.IsZero(m.Completed) { // not required
		return nil
	}

	if err := validate.FormatOf("completed", "body", "date-time", m.Completed.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateInterval(formats strfmt.Registry) error {
	if swag.IsZero(m.Interval) { // not required
		return nil
	}

	if err := validate.MinimumInt("interval", "body", *m.Interval, 1, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("interval", "body", *m.Interval, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateJobID(formats strfmt.Registry) error {

	if err := validate.Required("job_id", "body", m.JobID); err != nil {
		return err
	}

	if err := validate.FormatOf("job_id", "body", "uuid", m.JobID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 255); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateScheduled(formats strfmt.Registry) error {
	if swag.IsZero(m.Scheduled) { // not required
		return nil
	}

	if err := validate.FormatOf("scheduled", "body", "date-time", m.Scheduled.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateStarted(formats strfmt.Registry) error {
	if swag.IsZero(m.Started) { // not required
		return nil
	}

	if err := validate.FormatOf("started", "body", "date-time", m.Started.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if m.Status != nil {
		if err := m.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobResult) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) validateUser(formats strfmt.Registry) error {
	if swag.IsZero(m.User) { // not required
		return nil
	}

	if m.User != nil {
		if err := m.User.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this job result based on the context it is used
func (m *JobResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateObjType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUser(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobResult) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", strfmt.DateTime(m.Created)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateObjType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "obj_type", "body", string(m.ObjType)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if m.Status != nil {
		if err := m.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("status")
			}
			return err
		}
	}

	return nil
}

func (m *JobResult) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

func (m *JobResult) contextValidateUser(ctx context.Context, formats strfmt.Registry) error {

	if m.User != nil {
		if err := m.User.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("user")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobResult) UnmarshalBinary(b []byte) error {
	var res JobResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// JobResultStatus Status
//
// swagger:model JobResultStatus
type JobResultStatus struct {

	// label
	// Required: true
	// Enum: [Pending Scheduled Running Completed Errored Failed]
	Label *string `json:"label"`

	// value
	// Required: true
	// Enum: [pending scheduled running completed errored failed]
	Value *string `json:"value"`
}

// Validate validates this job result status
func (m *JobResultStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var jobResultStatusTypeLabelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Pending","Scheduled","Running","Completed","Errored","Failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobResultStatusTypeLabelPropEnum = append(jobResultStatusTypeLabelPropEnum, v)
	}
}

const (

	// JobResultStatusLabelPending captures enum value "Pending"
	JobResultStatusLabelPending string = "Pending"

	// JobResultStatusLabelScheduled captures enum value "Scheduled"
	JobResultStatusLabelScheduled string = "Scheduled"

	// JobResultStatusLabelRunning captures enum value "Running"
	JobResultStatusLabelRunning string = "Running"

	// JobResultStatusLabelCompleted captures enum value "Completed"
	JobResultStatusLabelCompleted string = "Completed"

	// JobResultStatusLabelErrored captures enum value "Errored"
	JobResultStatusLabelErrored string = "Errored"

	// JobResultStatusLabelFailed captures enum value "Failed"
	JobResultStatusLabelFailed string = "Failed"
)

// prop value enum
func (m *JobResultStatus) validateLabelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jobResultStatusTypeLabelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JobResultStatus) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"label", "body", m.Label); err != nil {
		return err
	}

	// value enum
	if err := m.validateLabelEnum("status"+"."+"label", "body", *m.Label); err != nil {
		return err
	}

	return nil
}

var jobResultStatusTypeValuePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","scheduled","running","completed","errored","failed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		jobResultStatusTypeValuePropEnum = append(jobResultStatusTypeValuePropEnum, v)
	}
}

const (

	// JobResultStatusValuePending captures enum value "pending"
	JobResultStatusValuePending string = "pending"

	// JobResultStatusValueScheduled captures enum value "scheduled"
	JobResultStatusValueScheduled string = "scheduled"

	// JobResultStatusValueRunning captures enum value "running"
	JobResultStatusValueRunning string = "running"

	// JobResultStatusValueCompleted captures enum value "completed"
	JobResultStatusValueCompleted string = "completed"

	// JobResultStatusValueErrored captures enum value "errored"
	JobResultStatusValueErrored string = "errored"

	// JobResultStatusValueFailed captures enum value "failed"
	JobResultStatusValueFailed string = "failed"
)

// prop value enum
func (m *JobResultStatus) validateValueEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, jobResultStatusTypeValuePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *JobResultStatus) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("status"+"."+"value", "body", m.Value); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueEnum("status"+"."+"value", "body", *m.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this job result status based on the context it is used
func (m *JobResultStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *JobResultStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobResultStatus) UnmarshalBinary(b []byte) error {
	var res JobResultStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
