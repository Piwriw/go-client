// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/piwriw/go-client/pkg/sdk/v2.0/models"
)

// CreateWebhookPolicyOfProjectReader is a Reader for the CreateWebhookPolicyOfProject structure.
type CreateWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateWebhookPolicyOfProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateWebhookPolicyOfProjectCreated creates a CreateWebhookPolicyOfProjectCreated with default headers values
func NewCreateWebhookPolicyOfProjectCreated() *CreateWebhookPolicyOfProjectCreated {
	return &CreateWebhookPolicyOfProjectCreated{}
}

/*
CreateWebhookPolicyOfProjectCreated describes a response with status code 201, with default header values.

Project webhook policy create successfully.
*/
type CreateWebhookPolicyOfProjectCreated struct {

	/* The location of the resource
	 */
	Location string

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this create webhook policy of project created response has a 2xx status code
func (o *CreateWebhookPolicyOfProjectCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create webhook policy of project created response has a 3xx status code
func (o *CreateWebhookPolicyOfProjectCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create webhook policy of project created response has a 4xx status code
func (o *CreateWebhookPolicyOfProjectCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create webhook policy of project created response has a 5xx status code
func (o *CreateWebhookPolicyOfProjectCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create webhook policy of project created response a status code equal to that given
func (o *CreateWebhookPolicyOfProjectCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateWebhookPolicyOfProjectCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectCreated ", 201)
}

func (o *CreateWebhookPolicyOfProjectCreated) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectCreated ", 201)
}

func (o *CreateWebhookPolicyOfProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Location
	hdrLocation := response.GetHeader("Location")

	if hdrLocation != "" {
		o.Location = hdrLocation
	}

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewCreateWebhookPolicyOfProjectBadRequest creates a CreateWebhookPolicyOfProjectBadRequest with default headers values
func NewCreateWebhookPolicyOfProjectBadRequest() *CreateWebhookPolicyOfProjectBadRequest {
	return &CreateWebhookPolicyOfProjectBadRequest{}
}

/*
CreateWebhookPolicyOfProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type CreateWebhookPolicyOfProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create webhook policy of project bad request response has a 2xx status code
func (o *CreateWebhookPolicyOfProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create webhook policy of project bad request response has a 3xx status code
func (o *CreateWebhookPolicyOfProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create webhook policy of project bad request response has a 4xx status code
func (o *CreateWebhookPolicyOfProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this create webhook policy of project bad request response has a 5xx status code
func (o *CreateWebhookPolicyOfProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this create webhook policy of project bad request response a status code equal to that given
func (o *CreateWebhookPolicyOfProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *CreateWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateWebhookPolicyOfProjectUnauthorized creates a CreateWebhookPolicyOfProjectUnauthorized with default headers values
func NewCreateWebhookPolicyOfProjectUnauthorized() *CreateWebhookPolicyOfProjectUnauthorized {
	return &CreateWebhookPolicyOfProjectUnauthorized{}
}

/*
CreateWebhookPolicyOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type CreateWebhookPolicyOfProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create webhook policy of project unauthorized response has a 2xx status code
func (o *CreateWebhookPolicyOfProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create webhook policy of project unauthorized response has a 3xx status code
func (o *CreateWebhookPolicyOfProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create webhook policy of project unauthorized response has a 4xx status code
func (o *CreateWebhookPolicyOfProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this create webhook policy of project unauthorized response has a 5xx status code
func (o *CreateWebhookPolicyOfProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this create webhook policy of project unauthorized response a status code equal to that given
func (o *CreateWebhookPolicyOfProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateWebhookPolicyOfProjectForbidden creates a CreateWebhookPolicyOfProjectForbidden with default headers values
func NewCreateWebhookPolicyOfProjectForbidden() *CreateWebhookPolicyOfProjectForbidden {
	return &CreateWebhookPolicyOfProjectForbidden{}
}

/*
CreateWebhookPolicyOfProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type CreateWebhookPolicyOfProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create webhook policy of project forbidden response has a 2xx status code
func (o *CreateWebhookPolicyOfProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create webhook policy of project forbidden response has a 3xx status code
func (o *CreateWebhookPolicyOfProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create webhook policy of project forbidden response has a 4xx status code
func (o *CreateWebhookPolicyOfProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this create webhook policy of project forbidden response has a 5xx status code
func (o *CreateWebhookPolicyOfProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this create webhook policy of project forbidden response a status code equal to that given
func (o *CreateWebhookPolicyOfProjectForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *CreateWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewCreateWebhookPolicyOfProjectInternalServerError creates a CreateWebhookPolicyOfProjectInternalServerError with default headers values
func NewCreateWebhookPolicyOfProjectInternalServerError() *CreateWebhookPolicyOfProjectInternalServerError {
	return &CreateWebhookPolicyOfProjectInternalServerError{}
}

/*
CreateWebhookPolicyOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type CreateWebhookPolicyOfProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this create webhook policy of project internal server error response has a 2xx status code
func (o *CreateWebhookPolicyOfProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this create webhook policy of project internal server error response has a 3xx status code
func (o *CreateWebhookPolicyOfProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create webhook policy of project internal server error response has a 4xx status code
func (o *CreateWebhookPolicyOfProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this create webhook policy of project internal server error response has a 5xx status code
func (o *CreateWebhookPolicyOfProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this create webhook policy of project internal server error response a status code equal to that given
func (o *CreateWebhookPolicyOfProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/webhook/policies][%d] createWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *CreateWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
