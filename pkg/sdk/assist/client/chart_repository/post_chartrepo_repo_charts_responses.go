// Code generated by go-swagger; DO NOT EDIT.

package chart_repository

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// PostChartrepoRepoChartsReader is a Reader for the PostChartrepoRepoCharts structure.
type PostChartrepoRepoChartsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostChartrepoRepoChartsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPostChartrepoRepoChartsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPostChartrepoRepoChartsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostChartrepoRepoChartsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostChartrepoRepoChartsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 507:
		result := NewPostChartrepoRepoChartsInsufficientStorage()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPostChartrepoRepoChartsCreated creates a PostChartrepoRepoChartsCreated with default headers values
func NewPostChartrepoRepoChartsCreated() *PostChartrepoRepoChartsCreated {
	return &PostChartrepoRepoChartsCreated{}
}

/*PostChartrepoRepoChartsCreated handles this case with default header values.

The specified chart is successfully uploaded.
*/
type PostChartrepoRepoChartsCreated struct {
}

func (o *PostChartrepoRepoChartsCreated) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsCreated ", 201)
}

func (o *PostChartrepoRepoChartsCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsUnauthorized creates a PostChartrepoRepoChartsUnauthorized with default headers values
func NewPostChartrepoRepoChartsUnauthorized() *PostChartrepoRepoChartsUnauthorized {
	return &PostChartrepoRepoChartsUnauthorized{}
}

/*PostChartrepoRepoChartsUnauthorized handles this case with default header values.

Unauthorized
*/
type PostChartrepoRepoChartsUnauthorized struct {
}

func (o *PostChartrepoRepoChartsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsUnauthorized ", 401)
}

func (o *PostChartrepoRepoChartsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsForbidden creates a PostChartrepoRepoChartsForbidden with default headers values
func NewPostChartrepoRepoChartsForbidden() *PostChartrepoRepoChartsForbidden {
	return &PostChartrepoRepoChartsForbidden{}
}

/*PostChartrepoRepoChartsForbidden handles this case with default header values.

Operation is forbidden or quota exceeded
*/
type PostChartrepoRepoChartsForbidden struct {
}

func (o *PostChartrepoRepoChartsForbidden) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsForbidden ", 403)
}

func (o *PostChartrepoRepoChartsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsInternalServerError creates a PostChartrepoRepoChartsInternalServerError with default headers values
func NewPostChartrepoRepoChartsInternalServerError() *PostChartrepoRepoChartsInternalServerError {
	return &PostChartrepoRepoChartsInternalServerError{}
}

/*PostChartrepoRepoChartsInternalServerError handles this case with default header values.

Internal server error occurred
*/
type PostChartrepoRepoChartsInternalServerError struct {
}

func (o *PostChartrepoRepoChartsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInternalServerError ", 500)
}

func (o *PostChartrepoRepoChartsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPostChartrepoRepoChartsInsufficientStorage creates a PostChartrepoRepoChartsInsufficientStorage with default headers values
func NewPostChartrepoRepoChartsInsufficientStorage() *PostChartrepoRepoChartsInsufficientStorage {
	return &PostChartrepoRepoChartsInsufficientStorage{}
}

/*PostChartrepoRepoChartsInsufficientStorage handles this case with default header values.

Insufficient storage
*/
type PostChartrepoRepoChartsInsufficientStorage struct {
}

func (o *PostChartrepoRepoChartsInsufficientStorage) Error() string {
	return fmt.Sprintf("[POST /chartrepo/{repo}/charts][%d] postChartrepoRepoChartsInsufficientStorage ", 507)
}

func (o *PostChartrepoRepoChartsInsufficientStorage) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
