// Code generated by go-swagger; DO NOT EDIT.

package configure

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/piwriw/go-client/pkg/sdk/v2.0/models"
)

// GetConfigurationsReader is a Reader for the GetConfigurations structure.
type GetConfigurationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetConfigurationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetConfigurationsOK creates a GetConfigurationsOK with default headers values
func NewGetConfigurationsOK() *GetConfigurationsOK {
	return &GetConfigurationsOK{}
}

/*
GetConfigurationsOK describes a response with status code 200, with default header values.

Get system configurations successfully. The response body is a map.
*/
type GetConfigurationsOK struct {
	Payload *models.ConfigurationsResponse
}

// IsSuccess returns true when this get configurations o k response has a 2xx status code
func (o *GetConfigurationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get configurations o k response has a 3xx status code
func (o *GetConfigurationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configurations o k response has a 4xx status code
func (o *GetConfigurationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configurations o k response has a 5xx status code
func (o *GetConfigurationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get configurations o k response a status code equal to that given
func (o *GetConfigurationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetConfigurationsOK) Error() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationsOK) String() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationsOK) GetPayload() *models.ConfigurationsResponse {
	return o.Payload
}

func (o *GetConfigurationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ConfigurationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationsUnauthorized creates a GetConfigurationsUnauthorized with default headers values
func NewGetConfigurationsUnauthorized() *GetConfigurationsUnauthorized {
	return &GetConfigurationsUnauthorized{}
}

/*
GetConfigurationsUnauthorized describes a response with status code 401, with default header values.

User need to log in first.ß
*/
type GetConfigurationsUnauthorized struct {
}

// IsSuccess returns true when this get configurations unauthorized response has a 2xx status code
func (o *GetConfigurationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configurations unauthorized response has a 3xx status code
func (o *GetConfigurationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configurations unauthorized response has a 4xx status code
func (o *GetConfigurationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configurations unauthorized response has a 5xx status code
func (o *GetConfigurationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get configurations unauthorized response a status code equal to that given
func (o *GetConfigurationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *GetConfigurationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsUnauthorized ", 401)
}

func (o *GetConfigurationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsUnauthorized ", 401)
}

func (o *GetConfigurationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationsForbidden creates a GetConfigurationsForbidden with default headers values
func NewGetConfigurationsForbidden() *GetConfigurationsForbidden {
	return &GetConfigurationsForbidden{}
}

/*
GetConfigurationsForbidden describes a response with status code 403, with default header values.

User does not have permission of admin role.
*/
type GetConfigurationsForbidden struct {
}

// IsSuccess returns true when this get configurations forbidden response has a 2xx status code
func (o *GetConfigurationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configurations forbidden response has a 3xx status code
func (o *GetConfigurationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configurations forbidden response has a 4xx status code
func (o *GetConfigurationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get configurations forbidden response has a 5xx status code
func (o *GetConfigurationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get configurations forbidden response a status code equal to that given
func (o *GetConfigurationsForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *GetConfigurationsForbidden) Error() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsForbidden ", 403)
}

func (o *GetConfigurationsForbidden) String() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsForbidden ", 403)
}

func (o *GetConfigurationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetConfigurationsInternalServerError creates a GetConfigurationsInternalServerError with default headers values
func NewGetConfigurationsInternalServerError() *GetConfigurationsInternalServerError {
	return &GetConfigurationsInternalServerError{}
}

/*
GetConfigurationsInternalServerError describes a response with status code 500, with default header values.

Unexpected internal errors.
*/
type GetConfigurationsInternalServerError struct {
}

// IsSuccess returns true when this get configurations internal server error response has a 2xx status code
func (o *GetConfigurationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get configurations internal server error response has a 3xx status code
func (o *GetConfigurationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get configurations internal server error response has a 4xx status code
func (o *GetConfigurationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get configurations internal server error response has a 5xx status code
func (o *GetConfigurationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get configurations internal server error response a status code equal to that given
func (o *GetConfigurationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *GetConfigurationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsInternalServerError ", 500)
}

func (o *GetConfigurationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /configurations][%d] getConfigurationsInternalServerError ", 500)
}

func (o *GetConfigurationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
