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
)

// NewIpamServiceTemplatesListParams creates a new IpamServiceTemplatesListParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpamServiceTemplatesListParams() *IpamServiceTemplatesListParams {
	return &IpamServiceTemplatesListParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpamServiceTemplatesListParamsWithTimeout creates a new IpamServiceTemplatesListParams object
// with the ability to set a timeout on a request.
func NewIpamServiceTemplatesListParamsWithTimeout(timeout time.Duration) *IpamServiceTemplatesListParams {
	return &IpamServiceTemplatesListParams{
		timeout: timeout,
	}
}

// NewIpamServiceTemplatesListParamsWithContext creates a new IpamServiceTemplatesListParams object
// with the ability to set a context for a request.
func NewIpamServiceTemplatesListParamsWithContext(ctx context.Context) *IpamServiceTemplatesListParams {
	return &IpamServiceTemplatesListParams{
		Context: ctx,
	}
}

// NewIpamServiceTemplatesListParamsWithHTTPClient creates a new IpamServiceTemplatesListParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpamServiceTemplatesListParamsWithHTTPClient(client *http.Client) *IpamServiceTemplatesListParams {
	return &IpamServiceTemplatesListParams{
		HTTPClient: client,
	}
}

/*
IpamServiceTemplatesListParams contains all the parameters to send to the API endpoint

	for the ipam service templates list operation.

	Typically these are written to a http.Request.
*/
type IpamServiceTemplatesListParams struct {

	// Created.
	Created *string

	// CreatedGt.
	CreatedGt *string

	// CreatedGte.
	CreatedGte *string

	// CreatedLt.
	CreatedLt *string

	// CreatedLte.
	CreatedLte *string

	// Createdn.
	Createdn *string

	// ID.
	ID *string

	// IDGt.
	IDGt *string

	// IDGte.
	IDGte *string

	// IDLt.
	IDLt *string

	// IDLte.
	IDLte *string

	// IDn.
	IDn *string

	// LastUpdated.
	LastUpdated *string

	// LastUpdatedGt.
	LastUpdatedGt *string

	// LastUpdatedGte.
	LastUpdatedGte *string

	// LastUpdatedLt.
	LastUpdatedLt *string

	// LastUpdatedLte.
	LastUpdatedLte *string

	// LastUpdatedn.
	LastUpdatedn *string

	/* Limit.

	   Number of results to return per page.
	*/
	Limit *int64

	// Name.
	Name *string

	// NameEmpty.
	NameEmpty *string

	// NameIc.
	NameIc *string

	// NameIe.
	NameIe *string

	// NameIew.
	NameIew *string

	// NameIsw.
	NameIsw *string

	// Namen.
	Namen *string

	// NameNic.
	NameNic *string

	// NameNie.
	NameNie *string

	// NameNiew.
	NameNiew *string

	// NameNisw.
	NameNisw *string

	/* Offset.

	   The initial index from which to return the results.
	*/
	Offset *int64

	/* Ordering.

	   Which field to use when ordering the results.
	*/
	Ordering *string

	// Port.
	Port *float64

	// Protocol.
	Protocol *string

	// Protocoln.
	Protocoln *string

	// Q.
	Q *string

	// Tag.
	Tag *string

	// Tagn.
	Tagn *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipam service templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesListParams) WithDefaults() *IpamServiceTemplatesListParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipam service templates list params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpamServiceTemplatesListParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithTimeout(timeout time.Duration) *IpamServiceTemplatesListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithContext(ctx context.Context) *IpamServiceTemplatesListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithHTTPClient(client *http.Client) *IpamServiceTemplatesListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCreated adds the created to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreated(created *string) *IpamServiceTemplatesListParams {
	o.SetCreated(created)
	return o
}

// SetCreated adds the created to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreated(created *string) {
	o.Created = created
}

// WithCreatedGt adds the createdGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreatedGt(createdGt *string) *IpamServiceTemplatesListParams {
	o.SetCreatedGt(createdGt)
	return o
}

// SetCreatedGt adds the createdGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreatedGt(createdGt *string) {
	o.CreatedGt = createdGt
}

// WithCreatedGte adds the createdGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreatedGte(createdGte *string) *IpamServiceTemplatesListParams {
	o.SetCreatedGte(createdGte)
	return o
}

// SetCreatedGte adds the createdGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreatedGte(createdGte *string) {
	o.CreatedGte = createdGte
}

// WithCreatedLt adds the createdLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreatedLt(createdLt *string) *IpamServiceTemplatesListParams {
	o.SetCreatedLt(createdLt)
	return o
}

// SetCreatedLt adds the createdLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreatedLt(createdLt *string) {
	o.CreatedLt = createdLt
}

// WithCreatedLte adds the createdLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreatedLte(createdLte *string) *IpamServiceTemplatesListParams {
	o.SetCreatedLte(createdLte)
	return o
}

// SetCreatedLte adds the createdLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreatedLte(createdLte *string) {
	o.CreatedLte = createdLte
}

// WithCreatedn adds the createdn to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithCreatedn(createdn *string) *IpamServiceTemplatesListParams {
	o.SetCreatedn(createdn)
	return o
}

// SetCreatedn adds the createdN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetCreatedn(createdn *string) {
	o.Createdn = createdn
}

// WithID adds the id to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithID(id *string) *IpamServiceTemplatesListParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetID(id *string) {
	o.ID = id
}

// WithIDGt adds the iDGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithIDGt(iDGt *string) *IpamServiceTemplatesListParams {
	o.SetIDGt(iDGt)
	return o
}

// SetIDGt adds the idGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetIDGt(iDGt *string) {
	o.IDGt = iDGt
}

// WithIDGte adds the iDGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithIDGte(iDGte *string) *IpamServiceTemplatesListParams {
	o.SetIDGte(iDGte)
	return o
}

// SetIDGte adds the idGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetIDGte(iDGte *string) {
	o.IDGte = iDGte
}

// WithIDLt adds the iDLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithIDLt(iDLt *string) *IpamServiceTemplatesListParams {
	o.SetIDLt(iDLt)
	return o
}

// SetIDLt adds the idLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetIDLt(iDLt *string) {
	o.IDLt = iDLt
}

// WithIDLte adds the iDLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithIDLte(iDLte *string) *IpamServiceTemplatesListParams {
	o.SetIDLte(iDLte)
	return o
}

// SetIDLte adds the idLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetIDLte(iDLte *string) {
	o.IDLte = iDLte
}

// WithIDn adds the iDn to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithIDn(iDn *string) *IpamServiceTemplatesListParams {
	o.SetIDn(iDn)
	return o
}

// SetIDn adds the idN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetIDn(iDn *string) {
	o.IDn = iDn
}

// WithLastUpdated adds the lastUpdated to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdated(lastUpdated *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdated(lastUpdated)
	return o
}

// SetLastUpdated adds the lastUpdated to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdated(lastUpdated *string) {
	o.LastUpdated = lastUpdated
}

// WithLastUpdatedGt adds the lastUpdatedGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdatedGt(lastUpdatedGt *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdatedGt(lastUpdatedGt)
	return o
}

// SetLastUpdatedGt adds the lastUpdatedGt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdatedGt(lastUpdatedGt *string) {
	o.LastUpdatedGt = lastUpdatedGt
}

// WithLastUpdatedGte adds the lastUpdatedGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdatedGte(lastUpdatedGte *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdatedGte(lastUpdatedGte)
	return o
}

// SetLastUpdatedGte adds the lastUpdatedGte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdatedGte(lastUpdatedGte *string) {
	o.LastUpdatedGte = lastUpdatedGte
}

// WithLastUpdatedLt adds the lastUpdatedLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdatedLt(lastUpdatedLt *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdatedLt(lastUpdatedLt)
	return o
}

// SetLastUpdatedLt adds the lastUpdatedLt to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdatedLt(lastUpdatedLt *string) {
	o.LastUpdatedLt = lastUpdatedLt
}

// WithLastUpdatedLte adds the lastUpdatedLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdatedLte(lastUpdatedLte *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdatedLte(lastUpdatedLte)
	return o
}

// SetLastUpdatedLte adds the lastUpdatedLte to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdatedLte(lastUpdatedLte *string) {
	o.LastUpdatedLte = lastUpdatedLte
}

// WithLastUpdatedn adds the lastUpdatedn to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLastUpdatedn(lastUpdatedn *string) *IpamServiceTemplatesListParams {
	o.SetLastUpdatedn(lastUpdatedn)
	return o
}

// SetLastUpdatedn adds the lastUpdatedN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLastUpdatedn(lastUpdatedn *string) {
	o.LastUpdatedn = lastUpdatedn
}

// WithLimit adds the limit to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithLimit(limit *int64) *IpamServiceTemplatesListParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithName adds the name to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithName(name *string) *IpamServiceTemplatesListParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetName(name *string) {
	o.Name = name
}

// WithNameEmpty adds the nameEmpty to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameEmpty(nameEmpty *string) *IpamServiceTemplatesListParams {
	o.SetNameEmpty(nameEmpty)
	return o
}

// SetNameEmpty adds the nameEmpty to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameEmpty(nameEmpty *string) {
	o.NameEmpty = nameEmpty
}

// WithNameIc adds the nameIc to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameIc(nameIc *string) *IpamServiceTemplatesListParams {
	o.SetNameIc(nameIc)
	return o
}

// SetNameIc adds the nameIc to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameIc(nameIc *string) {
	o.NameIc = nameIc
}

// WithNameIe adds the nameIe to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameIe(nameIe *string) *IpamServiceTemplatesListParams {
	o.SetNameIe(nameIe)
	return o
}

// SetNameIe adds the nameIe to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameIe(nameIe *string) {
	o.NameIe = nameIe
}

// WithNameIew adds the nameIew to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameIew(nameIew *string) *IpamServiceTemplatesListParams {
	o.SetNameIew(nameIew)
	return o
}

// SetNameIew adds the nameIew to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameIew(nameIew *string) {
	o.NameIew = nameIew
}

// WithNameIsw adds the nameIsw to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameIsw(nameIsw *string) *IpamServiceTemplatesListParams {
	o.SetNameIsw(nameIsw)
	return o
}

// SetNameIsw adds the nameIsw to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameIsw(nameIsw *string) {
	o.NameIsw = nameIsw
}

// WithNamen adds the namen to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNamen(namen *string) *IpamServiceTemplatesListParams {
	o.SetNamen(namen)
	return o
}

// SetNamen adds the nameN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNamen(namen *string) {
	o.Namen = namen
}

// WithNameNic adds the nameNic to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameNic(nameNic *string) *IpamServiceTemplatesListParams {
	o.SetNameNic(nameNic)
	return o
}

// SetNameNic adds the nameNic to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameNic(nameNic *string) {
	o.NameNic = nameNic
}

// WithNameNie adds the nameNie to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameNie(nameNie *string) *IpamServiceTemplatesListParams {
	o.SetNameNie(nameNie)
	return o
}

// SetNameNie adds the nameNie to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameNie(nameNie *string) {
	o.NameNie = nameNie
}

// WithNameNiew adds the nameNiew to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameNiew(nameNiew *string) *IpamServiceTemplatesListParams {
	o.SetNameNiew(nameNiew)
	return o
}

// SetNameNiew adds the nameNiew to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameNiew(nameNiew *string) {
	o.NameNiew = nameNiew
}

// WithNameNisw adds the nameNisw to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithNameNisw(nameNisw *string) *IpamServiceTemplatesListParams {
	o.SetNameNisw(nameNisw)
	return o
}

// SetNameNisw adds the nameNisw to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetNameNisw(nameNisw *string) {
	o.NameNisw = nameNisw
}

// WithOffset adds the offset to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithOffset(offset *int64) *IpamServiceTemplatesListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetOffset(offset *int64) {
	o.Offset = offset
}

// WithOrdering adds the ordering to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithOrdering(ordering *string) *IpamServiceTemplatesListParams {
	o.SetOrdering(ordering)
	return o
}

// SetOrdering adds the ordering to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetOrdering(ordering *string) {
	o.Ordering = ordering
}

// WithPort adds the port to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithPort(port *float64) *IpamServiceTemplatesListParams {
	o.SetPort(port)
	return o
}

// SetPort adds the port to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetPort(port *float64) {
	o.Port = port
}

// WithProtocol adds the protocol to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithProtocol(protocol *string) *IpamServiceTemplatesListParams {
	o.SetProtocol(protocol)
	return o
}

// SetProtocol adds the protocol to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetProtocol(protocol *string) {
	o.Protocol = protocol
}

// WithProtocoln adds the protocoln to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithProtocoln(protocoln *string) *IpamServiceTemplatesListParams {
	o.SetProtocoln(protocoln)
	return o
}

// SetProtocoln adds the protocolN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetProtocoln(protocoln *string) {
	o.Protocoln = protocoln
}

// WithQ adds the q to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithQ(q *string) *IpamServiceTemplatesListParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetQ(q *string) {
	o.Q = q
}

// WithTag adds the tag to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithTag(tag *string) *IpamServiceTemplatesListParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetTag(tag *string) {
	o.Tag = tag
}

// WithTagn adds the tagn to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) WithTagn(tagn *string) *IpamServiceTemplatesListParams {
	o.SetTagn(tagn)
	return o
}

// SetTagn adds the tagN to the ipam service templates list params
func (o *IpamServiceTemplatesListParams) SetTagn(tagn *string) {
	o.Tagn = tagn
}

// WriteToRequest writes these params to a swagger request
func (o *IpamServiceTemplatesListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Created != nil {

		// query param created
		var qrCreated string

		if o.Created != nil {
			qrCreated = *o.Created
		}
		qCreated := qrCreated
		if qCreated != "" {

			if err := r.SetQueryParam("created", qCreated); err != nil {
				return err
			}
		}
	}

	if o.CreatedGt != nil {

		// query param created__gt
		var qrCreatedGt string

		if o.CreatedGt != nil {
			qrCreatedGt = *o.CreatedGt
		}
		qCreatedGt := qrCreatedGt
		if qCreatedGt != "" {

			if err := r.SetQueryParam("created__gt", qCreatedGt); err != nil {
				return err
			}
		}
	}

	if o.CreatedGte != nil {

		// query param created__gte
		var qrCreatedGte string

		if o.CreatedGte != nil {
			qrCreatedGte = *o.CreatedGte
		}
		qCreatedGte := qrCreatedGte
		if qCreatedGte != "" {

			if err := r.SetQueryParam("created__gte", qCreatedGte); err != nil {
				return err
			}
		}
	}

	if o.CreatedLt != nil {

		// query param created__lt
		var qrCreatedLt string

		if o.CreatedLt != nil {
			qrCreatedLt = *o.CreatedLt
		}
		qCreatedLt := qrCreatedLt
		if qCreatedLt != "" {

			if err := r.SetQueryParam("created__lt", qCreatedLt); err != nil {
				return err
			}
		}
	}

	if o.CreatedLte != nil {

		// query param created__lte
		var qrCreatedLte string

		if o.CreatedLte != nil {
			qrCreatedLte = *o.CreatedLte
		}
		qCreatedLte := qrCreatedLte
		if qCreatedLte != "" {

			if err := r.SetQueryParam("created__lte", qCreatedLte); err != nil {
				return err
			}
		}
	}

	if o.Createdn != nil {

		// query param created__n
		var qrCreatedn string

		if o.Createdn != nil {
			qrCreatedn = *o.Createdn
		}
		qCreatedn := qrCreatedn
		if qCreatedn != "" {

			if err := r.SetQueryParam("created__n", qCreatedn); err != nil {
				return err
			}
		}
	}

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.IDGt != nil {

		// query param id__gt
		var qrIDGt string

		if o.IDGt != nil {
			qrIDGt = *o.IDGt
		}
		qIDGt := qrIDGt
		if qIDGt != "" {

			if err := r.SetQueryParam("id__gt", qIDGt); err != nil {
				return err
			}
		}
	}

	if o.IDGte != nil {

		// query param id__gte
		var qrIDGte string

		if o.IDGte != nil {
			qrIDGte = *o.IDGte
		}
		qIDGte := qrIDGte
		if qIDGte != "" {

			if err := r.SetQueryParam("id__gte", qIDGte); err != nil {
				return err
			}
		}
	}

	if o.IDLt != nil {

		// query param id__lt
		var qrIDLt string

		if o.IDLt != nil {
			qrIDLt = *o.IDLt
		}
		qIDLt := qrIDLt
		if qIDLt != "" {

			if err := r.SetQueryParam("id__lt", qIDLt); err != nil {
				return err
			}
		}
	}

	if o.IDLte != nil {

		// query param id__lte
		var qrIDLte string

		if o.IDLte != nil {
			qrIDLte = *o.IDLte
		}
		qIDLte := qrIDLte
		if qIDLte != "" {

			if err := r.SetQueryParam("id__lte", qIDLte); err != nil {
				return err
			}
		}
	}

	if o.IDn != nil {

		// query param id__n
		var qrIDn string

		if o.IDn != nil {
			qrIDn = *o.IDn
		}
		qIDn := qrIDn
		if qIDn != "" {

			if err := r.SetQueryParam("id__n", qIDn); err != nil {
				return err
			}
		}
	}

	if o.LastUpdated != nil {

		// query param last_updated
		var qrLastUpdated string

		if o.LastUpdated != nil {
			qrLastUpdated = *o.LastUpdated
		}
		qLastUpdated := qrLastUpdated
		if qLastUpdated != "" {

			if err := r.SetQueryParam("last_updated", qLastUpdated); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGt != nil {

		// query param last_updated__gt
		var qrLastUpdatedGt string

		if o.LastUpdatedGt != nil {
			qrLastUpdatedGt = *o.LastUpdatedGt
		}
		qLastUpdatedGt := qrLastUpdatedGt
		if qLastUpdatedGt != "" {

			if err := r.SetQueryParam("last_updated__gt", qLastUpdatedGt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedGte != nil {

		// query param last_updated__gte
		var qrLastUpdatedGte string

		if o.LastUpdatedGte != nil {
			qrLastUpdatedGte = *o.LastUpdatedGte
		}
		qLastUpdatedGte := qrLastUpdatedGte
		if qLastUpdatedGte != "" {

			if err := r.SetQueryParam("last_updated__gte", qLastUpdatedGte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLt != nil {

		// query param last_updated__lt
		var qrLastUpdatedLt string

		if o.LastUpdatedLt != nil {
			qrLastUpdatedLt = *o.LastUpdatedLt
		}
		qLastUpdatedLt := qrLastUpdatedLt
		if qLastUpdatedLt != "" {

			if err := r.SetQueryParam("last_updated__lt", qLastUpdatedLt); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedLte != nil {

		// query param last_updated__lte
		var qrLastUpdatedLte string

		if o.LastUpdatedLte != nil {
			qrLastUpdatedLte = *o.LastUpdatedLte
		}
		qLastUpdatedLte := qrLastUpdatedLte
		if qLastUpdatedLte != "" {

			if err := r.SetQueryParam("last_updated__lte", qLastUpdatedLte); err != nil {
				return err
			}
		}
	}

	if o.LastUpdatedn != nil {

		// query param last_updated__n
		var qrLastUpdatedn string

		if o.LastUpdatedn != nil {
			qrLastUpdatedn = *o.LastUpdatedn
		}
		qLastUpdatedn := qrLastUpdatedn
		if qLastUpdatedn != "" {

			if err := r.SetQueryParam("last_updated__n", qLastUpdatedn); err != nil {
				return err
			}
		}
	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64

		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {

			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameEmpty != nil {

		// query param name__empty
		var qrNameEmpty string

		if o.NameEmpty != nil {
			qrNameEmpty = *o.NameEmpty
		}
		qNameEmpty := qrNameEmpty
		if qNameEmpty != "" {

			if err := r.SetQueryParam("name__empty", qNameEmpty); err != nil {
				return err
			}
		}
	}

	if o.NameIc != nil {

		// query param name__ic
		var qrNameIc string

		if o.NameIc != nil {
			qrNameIc = *o.NameIc
		}
		qNameIc := qrNameIc
		if qNameIc != "" {

			if err := r.SetQueryParam("name__ic", qNameIc); err != nil {
				return err
			}
		}
	}

	if o.NameIe != nil {

		// query param name__ie
		var qrNameIe string

		if o.NameIe != nil {
			qrNameIe = *o.NameIe
		}
		qNameIe := qrNameIe
		if qNameIe != "" {

			if err := r.SetQueryParam("name__ie", qNameIe); err != nil {
				return err
			}
		}
	}

	if o.NameIew != nil {

		// query param name__iew
		var qrNameIew string

		if o.NameIew != nil {
			qrNameIew = *o.NameIew
		}
		qNameIew := qrNameIew
		if qNameIew != "" {

			if err := r.SetQueryParam("name__iew", qNameIew); err != nil {
				return err
			}
		}
	}

	if o.NameIsw != nil {

		// query param name__isw
		var qrNameIsw string

		if o.NameIsw != nil {
			qrNameIsw = *o.NameIsw
		}
		qNameIsw := qrNameIsw
		if qNameIsw != "" {

			if err := r.SetQueryParam("name__isw", qNameIsw); err != nil {
				return err
			}
		}
	}

	if o.Namen != nil {

		// query param name__n
		var qrNamen string

		if o.Namen != nil {
			qrNamen = *o.Namen
		}
		qNamen := qrNamen
		if qNamen != "" {

			if err := r.SetQueryParam("name__n", qNamen); err != nil {
				return err
			}
		}
	}

	if o.NameNic != nil {

		// query param name__nic
		var qrNameNic string

		if o.NameNic != nil {
			qrNameNic = *o.NameNic
		}
		qNameNic := qrNameNic
		if qNameNic != "" {

			if err := r.SetQueryParam("name__nic", qNameNic); err != nil {
				return err
			}
		}
	}

	if o.NameNie != nil {

		// query param name__nie
		var qrNameNie string

		if o.NameNie != nil {
			qrNameNie = *o.NameNie
		}
		qNameNie := qrNameNie
		if qNameNie != "" {

			if err := r.SetQueryParam("name__nie", qNameNie); err != nil {
				return err
			}
		}
	}

	if o.NameNiew != nil {

		// query param name__niew
		var qrNameNiew string

		if o.NameNiew != nil {
			qrNameNiew = *o.NameNiew
		}
		qNameNiew := qrNameNiew
		if qNameNiew != "" {

			if err := r.SetQueryParam("name__niew", qNameNiew); err != nil {
				return err
			}
		}
	}

	if o.NameNisw != nil {

		// query param name__nisw
		var qrNameNisw string

		if o.NameNisw != nil {
			qrNameNisw = *o.NameNisw
		}
		qNameNisw := qrNameNisw
		if qNameNisw != "" {

			if err := r.SetQueryParam("name__nisw", qNameNisw); err != nil {
				return err
			}
		}
	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int64

		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt64(qrOffset)
		if qOffset != "" {

			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}
	}

	if o.Ordering != nil {

		// query param ordering
		var qrOrdering string

		if o.Ordering != nil {
			qrOrdering = *o.Ordering
		}
		qOrdering := qrOrdering
		if qOrdering != "" {

			if err := r.SetQueryParam("ordering", qOrdering); err != nil {
				return err
			}
		}
	}

	if o.Port != nil {

		// query param port
		var qrPort float64

		if o.Port != nil {
			qrPort = *o.Port
		}
		qPort := swag.FormatFloat64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
				return err
			}
		}
	}

	if o.Protocol != nil {

		// query param protocol
		var qrProtocol string

		if o.Protocol != nil {
			qrProtocol = *o.Protocol
		}
		qProtocol := qrProtocol
		if qProtocol != "" {

			if err := r.SetQueryParam("protocol", qProtocol); err != nil {
				return err
			}
		}
	}

	if o.Protocoln != nil {

		// query param protocol__n
		var qrProtocoln string

		if o.Protocoln != nil {
			qrProtocoln = *o.Protocoln
		}
		qProtocoln := qrProtocoln
		if qProtocoln != "" {

			if err := r.SetQueryParam("protocol__n", qProtocoln); err != nil {
				return err
			}
		}
	}

	if o.Q != nil {

		// query param q
		var qrQ string

		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {

			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}
	}

	if o.Tag != nil {

		// query param tag
		var qrTag string

		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {

			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}
	}

	if o.Tagn != nil {

		// query param tag__n
		var qrTagn string

		if o.Tagn != nil {
			qrTagn = *o.Tagn
		}
		qTagn := qrTagn
		if qTagn != "" {

			if err := r.SetQueryParam("tag__n", qTagn); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
