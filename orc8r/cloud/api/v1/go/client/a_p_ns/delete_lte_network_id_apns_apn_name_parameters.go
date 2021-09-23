// Code generated by go-swagger; DO NOT EDIT.

package a_p_ns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteLTENetworkIDAPNSAPNNameParams creates a new DeleteLTENetworkIDAPNSAPNNameParams object
// with the default values initialized.
func NewDeleteLTENetworkIDAPNSAPNNameParams() *DeleteLTENetworkIDAPNSAPNNameParams {
	var ()
	return &DeleteLTENetworkIDAPNSAPNNameParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteLTENetworkIDAPNSAPNNameParamsWithTimeout creates a new DeleteLTENetworkIDAPNSAPNNameParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteLTENetworkIDAPNSAPNNameParamsWithTimeout(timeout time.Duration) *DeleteLTENetworkIDAPNSAPNNameParams {
	var ()
	return &DeleteLTENetworkIDAPNSAPNNameParams{

		timeout: timeout,
	}
}

// NewDeleteLTENetworkIDAPNSAPNNameParamsWithContext creates a new DeleteLTENetworkIDAPNSAPNNameParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteLTENetworkIDAPNSAPNNameParamsWithContext(ctx context.Context) *DeleteLTENetworkIDAPNSAPNNameParams {
	var ()
	return &DeleteLTENetworkIDAPNSAPNNameParams{

		Context: ctx,
	}
}

// NewDeleteLTENetworkIDAPNSAPNNameParamsWithHTTPClient creates a new DeleteLTENetworkIDAPNSAPNNameParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteLTENetworkIDAPNSAPNNameParamsWithHTTPClient(client *http.Client) *DeleteLTENetworkIDAPNSAPNNameParams {
	var ()
	return &DeleteLTENetworkIDAPNSAPNNameParams{
		HTTPClient: client,
	}
}

/*DeleteLTENetworkIDAPNSAPNNameParams contains all the parameters to send to the API endpoint
for the delete LTE network ID APNS APN name operation typically these are written to a http.Request
*/
type DeleteLTENetworkIDAPNSAPNNameParams struct {

	/*APNName
	  Access Point Name

	*/
	APNName string
	/*NetworkID
	  Network ID

	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WithTimeout(timeout time.Duration) *DeleteLTENetworkIDAPNSAPNNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WithContext(ctx context.Context) *DeleteLTENetworkIDAPNSAPNNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WithHTTPClient(client *http.Client) *DeleteLTENetworkIDAPNSAPNNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAPNName adds the aPNName to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WithAPNName(aPNName string) *DeleteLTENetworkIDAPNSAPNNameParams {
	o.SetAPNName(aPNName)
	return o
}

// SetAPNName adds the apnName to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) SetAPNName(aPNName string) {
	o.APNName = aPNName
}

// WithNetworkID adds the networkID to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WithNetworkID(networkID string) *DeleteLTENetworkIDAPNSAPNNameParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the delete LTE network ID APNS APN name params
func (o *DeleteLTENetworkIDAPNSAPNNameParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteLTENetworkIDAPNSAPNNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param apn_name
	if err := r.SetPathParam("apn_name", o.APNName); err != nil {
		return err
	}

	// path param network_id
	if err := r.SetPathParam("network_id", o.NetworkID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}