// Code generated by go-swagger; DO NOT EDIT.

package retention

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// CreateRetentionReader is a Reader for the CreateRetention structure.
type CreateRetentionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateRetentionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateRetentionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateRetentionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateRetentionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateRetentionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateRetentionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateRetentionCreated creates a CreateRetentionCreated with default headers values
func NewCreateRetentionCreated() *CreateRetentionCreated {
	return &CreateRetentionCreated{}
}

/*CreateRetentionCreated handles this case with default header values.

Created
*/
type CreateRetentionCreated struct {
	/*The location of the resource
	 */
	Location string
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *CreateRetentionCreated) Error() string {
	return fmt.Sprintf("[POST /retentions][%d] createRetentionCreated ", 201)
}

func (o *CreateRetentionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewCreateRetentionBadRequest creates a CreateRetentionBadRequest with default headers values
func NewCreateRetentionBadRequest() *CreateRetentionBadRequest {
	return &CreateRetentionBadRequest{}
}

/*CreateRetentionBadRequest handles this case with default header values.

Bad request
*/
type CreateRetentionBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRetentionBadRequest) Error() string {
	return fmt.Sprintf("[POST /retentions][%d] createRetentionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateRetentionBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRetentionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRetentionUnauthorized creates a CreateRetentionUnauthorized with default headers values
func NewCreateRetentionUnauthorized() *CreateRetentionUnauthorized {
	return &CreateRetentionUnauthorized{}
}

/*CreateRetentionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateRetentionUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRetentionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /retentions][%d] createRetentionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateRetentionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRetentionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRetentionForbidden creates a CreateRetentionForbidden with default headers values
func NewCreateRetentionForbidden() *CreateRetentionForbidden {
	return &CreateRetentionForbidden{}
}

/*CreateRetentionForbidden handles this case with default header values.

Forbidden
*/
type CreateRetentionForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRetentionForbidden) Error() string {
	return fmt.Sprintf("[POST /retentions][%d] createRetentionForbidden  %+v", 403, o.Payload)
}

func (o *CreateRetentionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRetentionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateRetentionInternalServerError creates a CreateRetentionInternalServerError with default headers values
func NewCreateRetentionInternalServerError() *CreateRetentionInternalServerError {
	return &CreateRetentionInternalServerError{}
}

/*CreateRetentionInternalServerError handles this case with default header values.

Internal server error
*/
type CreateRetentionInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *CreateRetentionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /retentions][%d] createRetentionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateRetentionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateRetentionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
