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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/nonstdout/go-netbox/v3/netbox/models"
)

// NewDcimModuleTypesBulkUpdateParams creates a new DcimModuleTypesBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModuleTypesBulkUpdateParams() *DcimModuleTypesBulkUpdateParams {
	return &DcimModuleTypesBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModuleTypesBulkUpdateParamsWithTimeout creates a new DcimModuleTypesBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModuleTypesBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimModuleTypesBulkUpdateParams {
	return &DcimModuleTypesBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModuleTypesBulkUpdateParamsWithContext creates a new DcimModuleTypesBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimModuleTypesBulkUpdateParamsWithContext(ctx context.Context) *DcimModuleTypesBulkUpdateParams {
	return &DcimModuleTypesBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimModuleTypesBulkUpdateParamsWithHTTPClient creates a new DcimModuleTypesBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModuleTypesBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimModuleTypesBulkUpdateParams {
	return &DcimModuleTypesBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimModuleTypesBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim module types bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimModuleTypesBulkUpdateParams struct {

	// Data.
	Data *models.WritableModuleType

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim module types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesBulkUpdateParams) WithDefaults() *DcimModuleTypesBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim module types bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModuleTypesBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimModuleTypesBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) WithContext(ctx context.Context) *DcimModuleTypesBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimModuleTypesBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) WithData(data *models.WritableModuleType) *DcimModuleTypesBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim module types bulk update params
func (o *DcimModuleTypesBulkUpdateParams) SetData(data *models.WritableModuleType) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModuleTypesBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
