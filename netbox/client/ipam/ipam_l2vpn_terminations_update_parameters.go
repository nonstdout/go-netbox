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

// NewIpamL2vpnTerminationsUpdateParams creates a new IpamL2vpnTerminationsUpdateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamL2vpnTerminationsUpdateParams() *IpamL2vpnTerminationsUpdateParams {
	return &IpamL2vpnTerminationsUpdateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamL2vpnTerminationsUpdateParamsWithTimeout creates a new IpamL2vpnTerminationsUpdateParams object
// with the ability to set a timeout on a request.
func NewIpamL2vpnTerminationsUpdateParamsWithTimeout(timeout time.Duration) *IpamL2vpnTerminationsUpdateParams {
	return &IpamL2vpnTerminationsUpdateParams{
		timeout: timeout,
	}
}

// NewIpamL2vpnTerminationsUpdateParamsWithContext creates a new IpamL2vpnTerminationsUpdateParams object
// with the ability to set a context for a request.
func NewIpamL2vpnTerminationsUpdateParamsWithContext(ctx context.Context) *IpamL2vpnTerminationsUpdateParams {
	return &IpamL2vpnTerminationsUpdateParams{
		Context: ctx,
	}
}

// NewIpamL2vpnTerminationsUpdateParamsWithHTTPClient creates a new IpamL2vpnTerminationsUpdateParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamL2vpnTerminationsUpdateParamsWithHTTPClient(client *http.Client) *IpamL2vpnTerminationsUpdateParams {
	return &IpamL2vpnTerminationsUpdateParams{
		HTTPClient: client,
	}
}

/*
IpamL2vpnTerminationsUpdateParams contains all the parameters to send to the API endpoint

	for the ipam l2vpn terminations update operation.

	Typically these are written to a http.Request.
*/
type IpamL2vpnTerminationsUpdateParams struct {

	// Data.
	Data *models.WritableL2VPNTermination

	/* ID.

	   A unique integer value identifying this L2VPN termination.
	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam l2vpn terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsUpdateParams) WithDefaults() *IpamL2vpnTerminationsUpdateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam l2vpn terminations update params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamL2vpnTerminationsUpdateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) WithTimeout(timeout time.Duration) *IpamL2vpnTerminationsUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) WithContext(ctx context.Context) *IpamL2vpnTerminationsUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) WithHTTPClient(client *http.Client) *IpamL2vpnTerminationsUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) WithData(data *models.WritableL2VPNTermination) *IpamL2vpnTerminationsUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) SetData(data *models.WritableL2VPNTermination) {
	o.Data = data
}

// WithID adds the id to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) WithID(id int64) *IpamL2vpnTerminationsUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam l2vpn terminations update params
func (o *IpamL2vpnTerminationsUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IpamL2vpnTerminationsUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
