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

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/smutel/go-netbox/v3/netbox/models"
)

// ExtrasCustomFieldsCreateReader is a Reader for the ExtrasCustomFieldsCreate structure.
type ExtrasCustomFieldsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasCustomFieldsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewExtrasCustomFieldsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewExtrasCustomFieldsCreateCreated creates a ExtrasCustomFieldsCreateCreated with default headers values
func NewExtrasCustomFieldsCreateCreated() *ExtrasCustomFieldsCreateCreated {
	return &ExtrasCustomFieldsCreateCreated{}
}

/* ExtrasCustomFieldsCreateCreated describes a response with status code 201, with default header values.

ExtrasCustomFieldsCreateCreated extras custom fields create created
*/
type ExtrasCustomFieldsCreateCreated struct {
	Payload *models.CustomField
}

func (o *ExtrasCustomFieldsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/custom-fields/][%d] extrasCustomFieldsCreateCreated  %+v", 201, o.Payload)
}
func (o *ExtrasCustomFieldsCreateCreated) GetPayload() *models.CustomField {
	return o.Payload
}

func (o *ExtrasCustomFieldsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CustomField)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
