// Code generated by go-swagger; DO NOT EDIT.

package scan_all

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
)

// NewStopScanAllParams creates a new StopScanAllParams object
// with the default values initialized.
func NewStopScanAllParams() *StopScanAllParams {
	var ()
	return &StopScanAllParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewStopScanAllParamsWithTimeout creates a new StopScanAllParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewStopScanAllParamsWithTimeout(timeout time.Duration) *StopScanAllParams {
	var ()
	return &StopScanAllParams{

		timeout: timeout,
	}
}

// NewStopScanAllParamsWithContext creates a new StopScanAllParams object
// with the default values initialized, and the ability to set a context for a request
func NewStopScanAllParamsWithContext(ctx context.Context) *StopScanAllParams {
	var ()
	return &StopScanAllParams{

		Context: ctx,
	}
}

// NewStopScanAllParamsWithHTTPClient creates a new StopScanAllParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewStopScanAllParamsWithHTTPClient(client *http.Client) *StopScanAllParams {
	var ()
	return &StopScanAllParams{
		HTTPClient: client,
	}
}

/*StopScanAllParams contains all the parameters to send to the API endpoint
for the stop scan all operation typically these are written to a http.Request
*/
type StopScanAllParams struct {

	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the stop scan all params
func (o *StopScanAllParams) WithTimeout(timeout time.Duration) *StopScanAllParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the stop scan all params
func (o *StopScanAllParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the stop scan all params
func (o *StopScanAllParams) WithContext(ctx context.Context) *StopScanAllParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the stop scan all params
func (o *StopScanAllParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the stop scan all params
func (o *StopScanAllParams) WithHTTPClient(client *http.Client) *StopScanAllParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the stop scan all params
func (o *StopScanAllParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the stop scan all params
func (o *StopScanAllParams) WithXRequestID(xRequestID *string) *StopScanAllParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the stop scan all params
func (o *StopScanAllParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *StopScanAllParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
