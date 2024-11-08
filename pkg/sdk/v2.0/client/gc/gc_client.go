// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the gc client
type API interface {
	/*
	   CreateGCSchedule creates a gc schedule

	   This endpoint is for update gc schedule.
	*/
	CreateGCSchedule(ctx context.Context, params *CreateGCScheduleParams) (*CreateGCScheduleCreated, error)
	/*
	   GetGC gets gc status

	   This endpoint let user get gc status filtered by specific ID.*/
	GetGC(ctx context.Context, params *GetGCParams) (*GetGCOK, error)
	/*
	   GetGCHistory gets gc results

	   This endpoint let user get gc execution history.*/
	GetGCHistory(ctx context.Context, params *GetGCHistoryParams) (*GetGCHistoryOK, error)
	/*
	   GetGCLog gets gc job log

	   This endpoint let user get gc job logs filtered by specific ID.*/
	GetGCLog(ctx context.Context, params *GetGCLogParams) (*GetGCLogOK, error)
	/*
	   GetGCSchedule gets gc s schedule

	   This endpoint is for get schedule of gc job.*/
	GetGCSchedule(ctx context.Context, params *GetGCScheduleParams) (*GetGCScheduleOK, error)
	/*
	   UpdateGCSchedule updates gc s schedule

	   This endpoint is for update gc schedule.
	*/
	UpdateGCSchedule(ctx context.Context, params *UpdateGCScheduleParams) (*UpdateGCScheduleOK, error)
}

// New creates a new gc API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for gc API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
CreateGCSchedule creates a gc schedule

This endpoint is for update gc schedule.
*/
func (a *Client) CreateGCSchedule(ctx context.Context, params *CreateGCScheduleParams) (*CreateGCScheduleCreated, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createGCSchedule",
		Method:             "POST",
		PathPattern:        "/system/gc/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateGCScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateGCScheduleCreated), nil

}

/*
GetGC gets gc status

This endpoint let user get gc status filtered by specific ID.
*/
func (a *Client) GetGC(ctx context.Context, params *GetGCParams) (*GetGCOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGC",
		Method:             "GET",
		PathPattern:        "/system/gc/{gc_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGCReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGCOK), nil

}

/*
GetGCHistory gets gc results

This endpoint let user get gc execution history.
*/
func (a *Client) GetGCHistory(ctx context.Context, params *GetGCHistoryParams) (*GetGCHistoryOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGCHistory",
		Method:             "GET",
		PathPattern:        "/system/gc",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGCHistoryReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGCHistoryOK), nil

}

/*
GetGCLog gets gc job log

This endpoint let user get gc job logs filtered by specific ID.
*/
func (a *Client) GetGCLog(ctx context.Context, params *GetGCLogParams) (*GetGCLogOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGCLog",
		Method:             "GET",
		PathPattern:        "/system/gc/{gc_id}/log",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGCLogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGCLogOK), nil

}

/*
GetGCSchedule gets gc s schedule

This endpoint is for get schedule of gc job.
*/
func (a *Client) GetGCSchedule(ctx context.Context, params *GetGCScheduleParams) (*GetGCScheduleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getGCSchedule",
		Method:             "GET",
		PathPattern:        "/system/gc/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetGCScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetGCScheduleOK), nil

}

/*
UpdateGCSchedule updates gc s schedule

This endpoint is for update gc schedule.
*/
func (a *Client) UpdateGCSchedule(ctx context.Context, params *UpdateGCScheduleParams) (*UpdateGCScheduleOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateGCSchedule",
		Method:             "PUT",
		PathPattern:        "/system/gc/schedule",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UpdateGCScheduleReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateGCScheduleOK), nil

}
