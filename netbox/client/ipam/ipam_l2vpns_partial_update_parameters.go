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

package ipam

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

	"github.com/nonstdout/go-netbox/v3/netbox/models"
)

// NewIpamL2vpnsPartialUpdateParams creates a new IpamL2vpnsPartialUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnsPartialUpdateParams() *IpamL2vpnsPartialUpdateParams {
	return &IpamL2vpnsPartialUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnsPartialUpdateParamsWithTimeout creates a new IpamL2vpnsPartialUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnsPartialUpdateParamsWithTimeout(timeout time.Duration) *IpamL2vpnsPartialUpdateParams {
	return &IpamL2vpnsPartialUpdateParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnsPartialUpdateParamsWithContext creates a new IpamL2vpnsPartialUpdateParams object
// with the ability to set a context for a request.
func NewIpamL2vpnsPartialUpdateParamsWithContext(ctx context.Context) *IpamL2vpnsPartialUpdateParams {
	return &IpamL2vpnsPartialUpdateParams{
		Context: ctx,
	}
}

// NewIpamL2vpnsPartialUpdateParamsWithHTTPClient creates a new IpamL2vpnsPartialUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnsPartialUpdateParamsWithHTTPClient(client *http.Client) *IpamL2vpnsPartialUpdateParams {
	return &IpamL2vpnsPartialUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnsPartialUpdateParams contains all the parameters to send to the API endpoint

	for the ipam l2vpns partial update operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnsPartialUpdateParams struct {

	// Data.
	Data *models.WritableL2VPN

	/* ID.

	   A unique integer value identifying this L2VPN.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpns partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsPartialUpdateParams) WithDefaults() *IpamL2vpnsPartialUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpns partial update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnsPartialUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) WithTimeout(timeout time.Duration) *IpamL2vpnsPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) WithContext(ctx context.Context) *IpamL2vpnsPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) WithHTTPClient(client *http.Client) *IpamL2vpnsPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) WithData(data *models.WritableL2VPN) *IpamL2vpnsPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) SetData(data *models.WritableL2VPN) {
	o.Data = data
}

// WithID adds the id to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) WithID(id int64) *IpamL2vpnsPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam l2vpns partial update params
func (o *IpamL2vpnsPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnsPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
