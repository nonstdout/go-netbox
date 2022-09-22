// Code generated by go-swagger; DO NOT EDIT.

// Copyright (c) 2020 Samuel Mutel <12967891+smutel@users.noreply.github.com>
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
	"github.com/go-openapi/swag"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// NewDcimModulesPartialUpdateParams creates a new DcimModulesPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDcimModulesPartialUpdateParams() *DcimModulesPartialUpdateParams {
	return &DcimModulesPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDcimModulesPartialUpdateParamsWithTimeout creates a new DcimModulesPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewDcimModulesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimModulesPartialUpdateParams {
	return &DcimModulesPartialUpdateParams{
		timeout: timeout,
	}
}

// NewDcimModulesPartialUpdateParamsWithContext creates a new DcimModulesPartialUpdateParams object
// with the ability to set a context for a request.
func NewDcimModulesPartialUpdateParamsWithContext(ctx context.Context) *DcimModulesPartialUpdateParams {
	return &DcimModulesPartialUpdateParams{
		Context: ctx,
	}
}

// NewDcimModulesPartialUpdateParamsWithHTTPClient creates a new DcimModulesPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewDcimModulesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimModulesPartialUpdateParams {
	return &DcimModulesPartialUpdateParams{
		HTTPClient: client,
	}
}

/* DcimModulesPartialUpdateParams contains all the parameters to send to the API endpoint
   for the dcim modules partial update operation.

   Typically these are written to a http.Request.
*/
type DcimModulesPartialUpdateParams struct {

	// Data.
	Data *models.WritableModule

	/* ID.

	   A unique integer value identifying this module.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the dcim modules partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesPartialUpdateParams) WithDefaults() *DcimModulesPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the dcim modules partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DcimModulesPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimModulesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) WithContext(ctx context.Context) *DcimModulesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimModulesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) WithData(data *models.WritableModule) *DcimModulesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) SetData(data *models.WritableModule) {
	o.Data = data
}

// WithID adds the id to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) WithID(id int64) *DcimModulesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim modules partial update params
func (o *DcimModulesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimModulesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
