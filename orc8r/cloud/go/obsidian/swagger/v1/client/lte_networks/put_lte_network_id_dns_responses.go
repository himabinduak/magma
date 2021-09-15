// Code generated by go-swagger; DO NOT EDIT.

package lte_networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "magma/orc8r/cloud/go/obsidian/swagger/v1/models"
)

// PutLTENetworkIDDNSReader is a Reader for the PutLTENetworkIDDNS structure.
type PutLTENetworkIDDNSReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutLTENetworkIDDNSReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewPutLTENetworkIDDNSNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPutLTENetworkIDDNSDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPutLTENetworkIDDNSNoContent creates a PutLTENetworkIDDNSNoContent with default headers values
func NewPutLTENetworkIDDNSNoContent() *PutLTENetworkIDDNSNoContent {
	return &PutLTENetworkIDDNSNoContent{}
}

/*PutLTENetworkIDDNSNoContent handles this case with default header values.

Success
*/
type PutLTENetworkIDDNSNoContent struct {
}

func (o *PutLTENetworkIDDNSNoContent) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/dns][%d] putLteNetworkIdDnsNoContent ", 204)
}

func (o *PutLTENetworkIDDNSNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPutLTENetworkIDDNSDefault creates a PutLTENetworkIDDNSDefault with default headers values
func NewPutLTENetworkIDDNSDefault(code int) *PutLTENetworkIDDNSDefault {
	return &PutLTENetworkIDDNSDefault{
		_statusCode: code,
	}
}

/*PutLTENetworkIDDNSDefault handles this case with default header values.

Unexpected Error
*/
type PutLTENetworkIDDNSDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the put LTE network ID DNS default response
func (o *PutLTENetworkIDDNSDefault) Code() int {
	return o._statusCode
}

func (o *PutLTENetworkIDDNSDefault) Error() string {
	return fmt.Sprintf("[PUT /lte/{network_id}/dns][%d] PutLTENetworkIDDNS default  %+v", o._statusCode, o.Payload)
}

func (o *PutLTENetworkIDDNSDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *PutLTENetworkIDDNSDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}