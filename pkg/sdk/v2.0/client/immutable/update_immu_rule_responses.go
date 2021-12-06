// Code generated by go-swagger; DO NOT EDIT.

package immutable

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// UpdateImmuRuleReader is a Reader for the UpdateImmuRule structure.
type UpdateImmuRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateImmuRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateImmuRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateImmuRuleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateImmuRuleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateImmuRuleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateImmuRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateImmuRuleOK creates a UpdateImmuRuleOK with default headers values
func NewUpdateImmuRuleOK() *UpdateImmuRuleOK {
	return &UpdateImmuRuleOK{}
}

/*UpdateImmuRuleOK handles this case with default header values.

Success
*/
type UpdateImmuRuleOK struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string
}

func (o *UpdateImmuRuleOK) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] updateImmuRuleOK ", 200)
}

func (o *UpdateImmuRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	return nil
}

// NewUpdateImmuRuleBadRequest creates a UpdateImmuRuleBadRequest with default headers values
func NewUpdateImmuRuleBadRequest() *UpdateImmuRuleBadRequest {
	return &UpdateImmuRuleBadRequest{}
}

/*UpdateImmuRuleBadRequest handles this case with default header values.

Bad request
*/
type UpdateImmuRuleBadRequest struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateImmuRuleBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] updateImmuRuleBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateImmuRuleBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateImmuRuleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImmuRuleUnauthorized creates a UpdateImmuRuleUnauthorized with default headers values
func NewUpdateImmuRuleUnauthorized() *UpdateImmuRuleUnauthorized {
	return &UpdateImmuRuleUnauthorized{}
}

/*UpdateImmuRuleUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateImmuRuleUnauthorized struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateImmuRuleUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] updateImmuRuleUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateImmuRuleUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateImmuRuleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImmuRuleForbidden creates a UpdateImmuRuleForbidden with default headers values
func NewUpdateImmuRuleForbidden() *UpdateImmuRuleForbidden {
	return &UpdateImmuRuleForbidden{}
}

/*UpdateImmuRuleForbidden handles this case with default header values.

Forbidden
*/
type UpdateImmuRuleForbidden struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateImmuRuleForbidden) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] updateImmuRuleForbidden  %+v", 403, o.Payload)
}

func (o *UpdateImmuRuleForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateImmuRuleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateImmuRuleInternalServerError creates a UpdateImmuRuleInternalServerError with default headers values
func NewUpdateImmuRuleInternalServerError() *UpdateImmuRuleInternalServerError {
	return &UpdateImmuRuleInternalServerError{}
}

/*UpdateImmuRuleInternalServerError handles this case with default header values.

Internal server error
*/
type UpdateImmuRuleInternalServerError struct {
	/*The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

func (o *UpdateImmuRuleInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projects/{project_name_or_id}/immutabletagrules/{immutable_rule_id}][%d] updateImmuRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateImmuRuleInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *UpdateImmuRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Request-Id
	o.XRequestID = response.GetHeader("X-Request-Id")

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
