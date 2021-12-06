// Code generated by go-swagger; DO NOT EDIT.

package robotv1

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
	"github.com/go-openapi/swag"

	"github.com/goharbor/go-client/pkg/sdk/v2.0/models"
)

// NewCreateRobotV1Params creates a new CreateRobotV1Params object
// with the default values initialized.
func NewCreateRobotV1Params() *CreateRobotV1Params {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &CreateRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateRobotV1ParamsWithTimeout creates a new CreateRobotV1Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateRobotV1ParamsWithTimeout(timeout time.Duration) *CreateRobotV1Params {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &CreateRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,

		timeout: timeout,
	}
}

// NewCreateRobotV1ParamsWithContext creates a new CreateRobotV1Params object
// with the default values initialized, and the ability to set a context for a request
func NewCreateRobotV1ParamsWithContext(ctx context.Context) *CreateRobotV1Params {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &CreateRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,

		Context: ctx,
	}
}

// NewCreateRobotV1ParamsWithHTTPClient creates a new CreateRobotV1Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateRobotV1ParamsWithHTTPClient(client *http.Client) *CreateRobotV1Params {
	var (
		xIsResourceNameDefault = bool(false)
	)
	return &CreateRobotV1Params{
		XIsResourceName: &xIsResourceNameDefault,
		HTTPClient:      client,
	}
}

/*CreateRobotV1Params contains all the parameters to send to the API endpoint
for the create robot v1 operation typically these are written to a http.Request
*/
type CreateRobotV1Params struct {

	/*XIsResourceName
	  The flag to indicate whether the parameter which supports both name and id in the path is the name of the resource. When the X-Is-Resource-Name is false and the parameter can be converted to an integer, the parameter will be as an id, otherwise, it will be as a name.

	*/
	XIsResourceName *bool
	/*XRequestID
	  An unique ID for the request

	*/
	XRequestID *string
	/*ProjectNameOrID
	  The name or id of the project

	*/
	ProjectNameOrID string
	/*Robot
	  The JSON object of a robot account.

	*/
	Robot *models.RobotCreateV1

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create robot v1 params
func (o *CreateRobotV1Params) WithTimeout(timeout time.Duration) *CreateRobotV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create robot v1 params
func (o *CreateRobotV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create robot v1 params
func (o *CreateRobotV1Params) WithContext(ctx context.Context) *CreateRobotV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create robot v1 params
func (o *CreateRobotV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create robot v1 params
func (o *CreateRobotV1Params) WithHTTPClient(client *http.Client) *CreateRobotV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create robot v1 params
func (o *CreateRobotV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXIsResourceName adds the xIsResourceName to the create robot v1 params
func (o *CreateRobotV1Params) WithXIsResourceName(xIsResourceName *bool) *CreateRobotV1Params {
	o.SetXIsResourceName(xIsResourceName)
	return o
}

// SetXIsResourceName adds the xIsResourceName to the create robot v1 params
func (o *CreateRobotV1Params) SetXIsResourceName(xIsResourceName *bool) {
	o.XIsResourceName = xIsResourceName
}

// WithXRequestID adds the xRequestID to the create robot v1 params
func (o *CreateRobotV1Params) WithXRequestID(xRequestID *string) *CreateRobotV1Params {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the create robot v1 params
func (o *CreateRobotV1Params) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithProjectNameOrID adds the projectNameOrID to the create robot v1 params
func (o *CreateRobotV1Params) WithProjectNameOrID(projectNameOrID string) *CreateRobotV1Params {
	o.SetProjectNameOrID(projectNameOrID)
	return o
}

// SetProjectNameOrID adds the projectNameOrId to the create robot v1 params
func (o *CreateRobotV1Params) SetProjectNameOrID(projectNameOrID string) {
	o.ProjectNameOrID = projectNameOrID
}

// WithRobot adds the robot to the create robot v1 params
func (o *CreateRobotV1Params) WithRobot(robot *models.RobotCreateV1) *CreateRobotV1Params {
	o.SetRobot(robot)
	return o
}

// SetRobot adds the robot to the create robot v1 params
func (o *CreateRobotV1Params) SetRobot(robot *models.RobotCreateV1) {
	o.Robot = robot
}

// WriteToRequest writes these params to a swagger request
func (o *CreateRobotV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XIsResourceName != nil {

		// header param X-Is-Resource-Name
		if err := r.SetHeaderParam("X-Is-Resource-Name", swag.FormatBool(*o.XIsResourceName)); err != nil {
			return err
		}

	}

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}

	}

	// path param project_name_or_id
	if err := r.SetPathParam("project_name_or_id", o.ProjectNameOrID); err != nil {
		return err
	}

	if o.Robot != nil {
		if err := r.SetBodyParam(o.Robot); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
