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

// NewDcimManufacturersBulkUpdateParams creates a new DcimManufacturersBulkUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimManufacturersBulkUpdateParams() *DcimManufacturersBulkUpdateParams {
	return &DcimManufacturersBulkUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimManufacturersBulkUpdateParamsWithTimeout creates a new DcimManufacturersBulkUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimManufacturersBulkUpdateParamsWithTimeout(timeout time.Duration) *DcimManufacturersBulkUpdateParams {
	return &DcimManufacturersBulkUpdateParams{
		timeout: timeout,
	}
}

// NewDcimManufacturersBulkUpdateParamsWithContext creates a new DcimManufacturersBulkUpdateParams object
// with the ability to set a context for a request.
func NewDcimManufacturersBulkUpdateParamsWithContext(ctx context.Context) *DcimManufacturersBulkUpdateParams {
	return &DcimManufacturersBulkUpdateParams{
		Context: ctx,
	}
}

// NewDcimManufacturersBulkUpdateParamsWithHTTPClient creates a new DcimManufacturersBulkUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimManufacturersBulkUpdateParamsWithHTTPClient(client *http.Client) *DcimManufacturersBulkUpdateParams {
	return &DcimManufacturersBulkUpdateParams{
		HTTPClient: client,
	}
}

/*
DcimManufacturersBulkUpdateParams contains all the parameters to send to the API endpoint

	for the dcim manufacturers bulk update operation.

	Typically these are written to a http.Request.
*/
type DcimManufacturersBulkUpdateParams struct {

	// Data.
	Data *models.Manufacturer

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim manufacturers bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersBulkUpdateParams) WithDefaults() *DcimManufacturersBulkUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim manufacturers bulk update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimManufacturersBulkUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) WithTimeout(timeout time.Duration) *DcimManufacturersBulkUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) WithContext(ctx context.Context) *DcimManufacturersBulkUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) WithHTTPClient(client *http.Client) *DcimManufacturersBulkUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) WithData(data *models.Manufacturer) *DcimManufacturersBulkUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim manufacturers bulk update params
func (o *DcimManufacturersBulkUpdateParams) SetData(data *models.Manufacturer) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimManufacturersBulkUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
