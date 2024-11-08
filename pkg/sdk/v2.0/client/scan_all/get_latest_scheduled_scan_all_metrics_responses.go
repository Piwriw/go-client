// Code generated by go-swagger; DO NOT EDIT.

package scan_all

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/piwriw/go-client/pkg/sdk/v2.0/models"
)

// GetLatestScheduledScanAllMetricsReader is a Reader for the GetLatestScheduledScanAllMetrics structure.
type GetLatestScheduledScanAllMetricsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLatestScheduledScanAllMetricsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetLatestScheduledScanAllMetricsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetLatestScheduledScanAllMetricsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetLatestScheduledScanAllMetricsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 412:
		result := NewGetLatestScheduledScanAllMetricsPreconditionFailed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetLatestScheduledScanAllMetricsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetLatestScheduledScanAllMetricsOK creates a GetLatestScheduledScanAllMetricsOK with default headers values
func NewGetLatestScheduledScanAllMetricsOK() *GetLatestScheduledScanAllMetricsOK {
	return &GetLatestScheduledScanAllMetricsOK{}
}

/*
GetLatestScheduledScanAllMetricsOK describes a response with status code 200, with default header values.

OK
*/
type GetLatestScheduledScanAllMetricsOK struct {
	Payload *models.Stats
}

// IsSuccess returns true when this get latest scheduled scan all metrics o k response has a 2xx status code
func (o *GetLatestScheduledScanAllMetricsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get latest scheduled scan all metrics o k response has a 3xx status code
func (o *GetLatestScheduledScanAllMetricsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scheduled scan all metrics o k response has a 4xx status code
func (o *GetLatestScheduledScanAllMetricsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest scheduled scan all metrics o k response has a 5xx status code
func (o *GetLatestScheduledScanAllMetricsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scheduled scan all metrics o k response a status code equal to that given
func (o *GetLatestScheduledScanAllMetricsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetLatestScheduledScanAllMetricsOK) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsOK) String() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsOK  %+v", 200, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsOK) GetPayload() *models.Stats {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Stats)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsUnauthorized creates a GetLatestScheduledScanAllMetricsUnauthorized with default headers values
func NewGetLatestScheduledScanAllMetricsUnauthorized() *GetLatestScheduledScanAllMetricsUnauthorized {
	return &GetLatestScheduledScanAllMetricsUnauthorized{}
}

/*
GetLatestScheduledScanAllMetricsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetLatestScheduledScanAllMetricsUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scheduled scan all metrics unauthorized response has a 2xx status code
func (o *GetLatestScheduledScanAllMetricsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scheduled scan all metrics unauthorized response has a 3xx status code
func (o *GetLatestScheduledScanAllMetricsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scheduled scan all metrics unauthorized response has a 4xx status code
func (o *GetLatestScheduledScanAllMetricsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scheduled scan all metrics unauthorized response has a 5xx status code
func (o *GetLatestScheduledScanAllMetricsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scheduled scan all metrics unauthorized response a status code equal to that given
func (o *GetLatestScheduledScanAllMetricsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) String() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsForbidden creates a GetLatestScheduledScanAllMetricsForbidden with default headers values
func NewGetLatestScheduledScanAllMetricsForbidden() *GetLatestScheduledScanAllMetricsForbidden {
	return &GetLatestScheduledScanAllMetricsForbidden{}
}

/*
GetLatestScheduledScanAllMetricsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetLatestScheduledScanAllMetricsForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scheduled scan all metrics forbidden response has a 2xx status code
func (o *GetLatestScheduledScanAllMetricsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scheduled scan all metrics forbidden response has a 3xx status code
func (o *GetLatestScheduledScanAllMetricsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scheduled scan all metrics forbidden response has a 4xx status code
func (o *GetLatestScheduledScanAllMetricsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scheduled scan all metrics forbidden response has a 5xx status code
func (o *GetLatestScheduledScanAllMetricsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scheduled scan all metrics forbidden response a status code equal to that given
func (o *GetLatestScheduledScanAllMetricsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetLatestScheduledScanAllMetricsForbidden) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsForbidden) String() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsForbidden  %+v", 403, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsPreconditionFailed creates a GetLatestScheduledScanAllMetricsPreconditionFailed with default headers values
func NewGetLatestScheduledScanAllMetricsPreconditionFailed() *GetLatestScheduledScanAllMetricsPreconditionFailed {
	return &GetLatestScheduledScanAllMetricsPreconditionFailed{}
}

/*
GetLatestScheduledScanAllMetricsPreconditionFailed describes a response with status code 412, with default header values.

Precondition failed
*/
type GetLatestScheduledScanAllMetricsPreconditionFailed struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scheduled scan all metrics precondition failed response has a 2xx status code
func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scheduled scan all metrics precondition failed response has a 3xx status code
func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scheduled scan all metrics precondition failed response has a 4xx status code
func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) IsClientError() bool {
	return true
}

// IsServerError returns true when this get latest scheduled scan all metrics precondition failed response has a 5xx status code
func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) IsServerError() bool {
	return false
}

// IsCode returns true when this get latest scheduled scan all metrics precondition failed response a status code equal to that given
func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) IsCode(code int) bool {
	return code == 412
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) String() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsPreconditionFailed  %+v", 412, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsPreconditionFailed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetLatestScheduledScanAllMetricsInternalServerError creates a GetLatestScheduledScanAllMetricsInternalServerError with default headers values
func NewGetLatestScheduledScanAllMetricsInternalServerError() *GetLatestScheduledScanAllMetricsInternalServerError {
	return &GetLatestScheduledScanAllMetricsInternalServerError{}
}

/*
GetLatestScheduledScanAllMetricsInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetLatestScheduledScanAllMetricsInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get latest scheduled scan all metrics internal server error response has a 2xx status code
func (o *GetLatestScheduledScanAllMetricsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get latest scheduled scan all metrics internal server error response has a 3xx status code
func (o *GetLatestScheduledScanAllMetricsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get latest scheduled scan all metrics internal server error response has a 4xx status code
func (o *GetLatestScheduledScanAllMetricsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get latest scheduled scan all metrics internal server error response has a 5xx status code
func (o *GetLatestScheduledScanAllMetricsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get latest scheduled scan all metrics internal server error response a status code equal to that given
func (o *GetLatestScheduledScanAllMetricsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) String() string {
	return fmt.Sprintf("[GET /scans/schedule/metrics][%d] getLatestScheduledScanAllMetricsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetLatestScheduledScanAllMetricsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
