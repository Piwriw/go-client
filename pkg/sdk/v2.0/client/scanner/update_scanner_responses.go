// Code generated by go-swagger; DO NOT EDIT.

package scanner

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// UpdateScannerReader is a Reader for the UpdateScanner structure.
type UpdateScannerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateScannerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateScannerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewUpdateScannerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateScannerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateScannerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateScannerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateScannerOK creates a UpdateScannerOK with default headers values
func NewUpdateScannerOK() *UpdateScannerOK {
	return &UpdateScannerOK{}
}

/*UpdateScannerOK handles this case with default header values.

Success
*/
type UpdateScannerOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateScannerOK) Error() string {
	return fmt.Sprintf("[PUT /scanners/{registration_id}][%d] updateScannerOK ", 200)
}

func (o *UpdateScannerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateScannerUnauthorized creates a UpdateScannerUnauthorized with default headers values
func NewUpdateScannerUnauthorized() *UpdateScannerUnauthorized {
	return &UpdateScannerUnauthorized{}
}

/*UpdateScannerUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateScannerUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateScannerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /scanners/{registration_id}][%d] updateScannerUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateScannerUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateScannerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScannerForbidden creates a UpdateScannerForbidden with default headers values
func NewUpdateScannerForbidden() *UpdateScannerForbidden {
	return &UpdateScannerForbidden{}
}

/*UpdateScannerForbidden handles this case with default header values.

Forbidden
*/
type UpdateScannerForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateScannerForbidden) Error() string {
	return fmt.Sprintf("[PUT /scanners/{registration_id}][%d] updateScannerForbidden  %+v", 403, o.Payload)
}

func (o *UpdateScannerForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateScannerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScannerNotFound creates a UpdateScannerNotFound with default headers values
func NewUpdateScannerNotFound() *UpdateScannerNotFound {
	return &UpdateScannerNotFound{}
}

/*UpdateScannerNotFound handles this case with default header values.

Not found
*/
type UpdateScannerNotFound struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateScannerNotFound) Error() string {
	return fmt.Sprintf("[PUT /scanners/{registration_id}][%d] updateScannerNotFound  %+v", 404, o.Payload)
}

func (o *UpdateScannerNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateScannerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateScannerInternalServerError creates a UpdateScannerInternalServerError with default headers values
func NewUpdateScannerInternalServerError() *UpdateScannerInternalServerError {
	return &UpdateScannerInternalServerError{}
}

/*UpdateScannerInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateScannerInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateScannerInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /scanners/{registration_id}][%d] updateScannerInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateScannerInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateScannerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
