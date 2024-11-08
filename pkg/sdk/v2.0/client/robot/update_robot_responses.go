// Code generated by go-swagger; DO NOT EDIT.

package robot

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/piwriw/go-client/pkg/sdk/v2.0/models"
)

// UpdateRobotReader is a Reader for the UpdateRobot structure.
type UpdateRobotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRobotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRobotOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateRobotBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateRobotUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateRobotForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateRobotNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateRobotConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateRobotInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateRobotOK creates a UpdateRobotOK with default headers values
func NewUpdateRobotOK() *UpdateRobotOK {
	return &UpdateRobotOK{}
}

/*
UpdateRobotOK describes a response with status code 200, with default header values.

Success
*/
type UpdateRobotOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this update robot o k response has a 2xx status code
func (o *UpdateRobotOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update robot o k response has a 3xx status code
func (o *UpdateRobotOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot o k response has a 4xx status code
func (o *UpdateRobotOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update robot o k response has a 5xx status code
func (o *UpdateRobotOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot o k response a status code equal to that given
func (o *UpdateRobotOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateRobotOK) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotOK ", 200)
}

func (o *UpdateRobotOK) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotOK ", 200)
}

func (o *UpdateRobotOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewUpdateRobotBadRequest creates a UpdateRobotBadRequest with default headers values
func NewUpdateRobotBadRequest() *UpdateRobotBadRequest {
	return &UpdateRobotBadRequest{}
}

/*
UpdateRobotBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type UpdateRobotBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot bad request response has a 2xx status code
func (o *UpdateRobotBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot bad request response has a 3xx status code
func (o *UpdateRobotBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot bad request response has a 4xx status code
func (o *UpdateRobotBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this update robot bad request response has a 5xx status code
func (o *UpdateRobotBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot bad request response a status code equal to that given
func (o *UpdateRobotBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *UpdateRobotBadRequest) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRobotBadRequest) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateRobotBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotUnauthorized creates a UpdateRobotUnauthorized with default headers values
func NewUpdateRobotUnauthorized() *UpdateRobotUnauthorized {
	return &UpdateRobotUnauthorized{}
}

/*
UpdateRobotUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type UpdateRobotUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot unauthorized response has a 2xx status code
func (o *UpdateRobotUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot unauthorized response has a 3xx status code
func (o *UpdateRobotUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot unauthorized response has a 4xx status code
func (o *UpdateRobotUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this update robot unauthorized response has a 5xx status code
func (o *UpdateRobotUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot unauthorized response a status code equal to that given
func (o *UpdateRobotUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *UpdateRobotUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRobotUnauthorized) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateRobotUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotForbidden creates a UpdateRobotForbidden with default headers values
func NewUpdateRobotForbidden() *UpdateRobotForbidden {
	return &UpdateRobotForbidden{}
}

/*
UpdateRobotForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type UpdateRobotForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot forbidden response has a 2xx status code
func (o *UpdateRobotForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot forbidden response has a 3xx status code
func (o *UpdateRobotForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot forbidden response has a 4xx status code
func (o *UpdateRobotForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this update robot forbidden response has a 5xx status code
func (o *UpdateRobotForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot forbidden response a status code equal to that given
func (o *UpdateRobotForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *UpdateRobotForbidden) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRobotForbidden) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotForbidden  %+v", 403, o.Payload)
}

func (o *UpdateRobotForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotNotFound creates a UpdateRobotNotFound with default headers values
func NewUpdateRobotNotFound() *UpdateRobotNotFound {
	return &UpdateRobotNotFound{}
}

/*
UpdateRobotNotFound describes a response with status code 404, with default header values.

Not found
*/
type UpdateRobotNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot not found response has a 2xx status code
func (o *UpdateRobotNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot not found response has a 3xx status code
func (o *UpdateRobotNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot not found response has a 4xx status code
func (o *UpdateRobotNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this update robot not found response has a 5xx status code
func (o *UpdateRobotNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot not found response a status code equal to that given
func (o *UpdateRobotNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *UpdateRobotNotFound) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRobotNotFound) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotNotFound  %+v", 404, o.Payload)
}

func (o *UpdateRobotNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotConflict creates a UpdateRobotConflict with default headers values
func NewUpdateRobotConflict() *UpdateRobotConflict {
	return &UpdateRobotConflict{}
}

/*
UpdateRobotConflict describes a response with status code 409, with default header values.

Conflict
*/
type UpdateRobotConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot conflict response has a 2xx status code
func (o *UpdateRobotConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot conflict response has a 3xx status code
func (o *UpdateRobotConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot conflict response has a 4xx status code
func (o *UpdateRobotConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this update robot conflict response has a 5xx status code
func (o *UpdateRobotConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this update robot conflict response a status code equal to that given
func (o *UpdateRobotConflict) IsCode(code int) bool {
	return code == 409
}

func (o *UpdateRobotConflict) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotConflict  %+v", 409, o.Payload)
}

func (o *UpdateRobotConflict) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotConflict  %+v", 409, o.Payload)
}

func (o *UpdateRobotConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewUpdateRobotInternalServerError creates a UpdateRobotInternalServerError with default headers values
func NewUpdateRobotInternalServerError() *UpdateRobotInternalServerError {
	return &UpdateRobotInternalServerError{}
}

/*
UpdateRobotInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type UpdateRobotInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this update robot internal server error response has a 2xx status code
func (o *UpdateRobotInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this update robot internal server error response has a 3xx status code
func (o *UpdateRobotInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update robot internal server error response has a 4xx status code
func (o *UpdateRobotInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this update robot internal server error response has a 5xx status code
func (o *UpdateRobotInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this update robot internal server error response a status code equal to that given
func (o *UpdateRobotInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *UpdateRobotInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRobotInternalServerError) String() string {
	return fmt.Sprintf("[PUT /robots/{robot_id}][%d] updateRobotInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRobotInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateRobotInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
