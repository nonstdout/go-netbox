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
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/nonstdout/go-netbox/v3/netbox/models"
)

// IpamVlanGroupsCreateReader is a Reader for the IpamVlanGroupsCreate structure.
type IpamVlanGroupsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamVlanGroupsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamVlanGroupsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamVlanGroupsCreateCreated creates a IpamVlanGroupsCreateCreated with default headers values
func NewIpamVlanGroupsCreateCreated() *IpamVlanGroupsCreateCreated {
	return &IpamVlanGroupsCreateCreated{}
}

/*
IpamVlanGroupsCreateCreated describes a response with status code 201, with default header values.

IpamVlanGroupsCreateCreated ipam vlan groups create created
*/
type IpamVlanGroupsCreateCreated struct {
	Payload *models.VLANGroup
}

// IsSuccess returns true when this ipam vlan groups create created response has a 2xx status code
func (o *IpamVlanGroupsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam vlan groups create created response has a 3xx status code
func (o *IpamVlanGroupsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam vlan groups create created response has a 4xx status code
func (o *IpamVlanGroupsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam vlan groups create created response has a 5xx status code
func (o *IpamVlanGroupsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam vlan groups create created response a status code equal to that given
func (o *IpamVlanGroupsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamVlanGroupsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipamVlanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/vlan-groups/][%d] ipamVlanGroupsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamVlanGroupsCreateCreated) GetPayload() *models.VLANGroup {
	return o.Payload
}

func (o *IpamVlanGroupsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VLANGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
