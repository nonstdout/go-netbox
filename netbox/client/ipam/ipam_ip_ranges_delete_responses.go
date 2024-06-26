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

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpamIPRangesDeleteReader is a Reader for the IpamIPRangesDelete structure.
type IpamIPRangesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamIPRangesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewIpamIPRangesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpamIPRangesDeleteNoContent creates a IpamIPRangesDeleteNoContent with default headers values
func NewIpamIPRangesDeleteNoContent() *IpamIPRangesDeleteNoContent {
	return &IpamIPRangesDeleteNoContent{}
}

/*
IpamIPRangesDeleteNoContent describes a response with status code 204, with default header values.

IpamIPRangesDeleteNoContent ipam Ip ranges delete no content
*/
type IpamIPRangesDeleteNoContent struct {
}

// IsSuccess returns true when this ipam Ip ranges delete no content response has a 2xx status code
func (o *IpamIPRangesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam Ip ranges delete no content response has a 3xx status code
func (o *IpamIPRangesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam Ip ranges delete no content response has a 4xx status code
func (o *IpamIPRangesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam Ip ranges delete no content response has a 5xx status code
func (o *IpamIPRangesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam Ip ranges delete no content response a status code equal to that given
func (o *IpamIPRangesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *IpamIPRangesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteNoContent ", 204)
}

func (o *IpamIPRangesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /ipam/ip-ranges/{id}/][%d] ipamIpRangesDeleteNoContent ", 204)
}

func (o *IpamIPRangesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
