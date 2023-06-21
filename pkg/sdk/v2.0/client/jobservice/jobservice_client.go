// Code generated by go-swagger; DO NOT EDIT.

package jobservice

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

//go:generate mockery -name API -inpkg

// API is the interface of the jobservice client
type API interface {
	/*
	   ActionGetJobLog gets job log by job id

	   Get job log by job id, it is only used by administrator*/
	ActionGetJobLog(ctx context.Context, params *ActionGetJobLogParams) (*ActionGetJobLogOK, error)
	/*
	   ActionPendingJobs stops and clean pause resume pending jobs in the queue

	   stop and clean, pause, resume pending jobs in the queue*/
	ActionPendingJobs(ctx context.Context, params *ActionPendingJobsParams) (*ActionPendingJobsOK, error)
	/*
	   GetWorkerPools gets worker pools

	   Get worker pools*/
	GetWorkerPools(ctx context.Context, params *GetWorkerPoolsParams) (*GetWorkerPoolsOK, error)
	/*
	   GetWorkers gets workers

	   Get workers in current pool*/
	GetWorkers(ctx context.Context, params *GetWorkersParams) (*GetWorkersOK, error)
	/*
	   ListJobQueues lists job queues

	   list job queue*/
	ListJobQueues(ctx context.Context, params *ListJobQueuesParams) (*ListJobQueuesOK, error)
	/*
	   StopRunningJob stops running job

	   Stop running job*/
	StopRunningJob(ctx context.Context, params *StopRunningJobParams) (*StopRunningJobOK, error)
}

// New creates a new jobservice API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry, authInfo runtime.ClientAuthInfoWriter) *Client {
	return &Client{
		transport: transport,
		formats:   formats,
		authInfo:  authInfo,
	}
}

/*
Client for jobservice API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
	authInfo  runtime.ClientAuthInfoWriter
}

/*
ActionGetJobLog gets job log by job id

Get job log by job id, it is only used by administrator
*/
func (a *Client) ActionGetJobLog(ctx context.Context, params *ActionGetJobLogParams) (*ActionGetJobLogOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actionGetJobLog",
		Method:             "GET",
		PathPattern:        "/jobservice/jobs/{job_id}/log",
		ProducesMediaTypes: []string{"text/plain"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ActionGetJobLogReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionGetJobLogOK), nil

}

/*
ActionPendingJobs stops and clean pause resume pending jobs in the queue

stop and clean, pause, resume pending jobs in the queue
*/
func (a *Client) ActionPendingJobs(ctx context.Context, params *ActionPendingJobsParams) (*ActionPendingJobsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "actionPendingJobs",
		Method:             "PUT",
		PathPattern:        "/jobservice/queues/{job_type}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ActionPendingJobsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ActionPendingJobsOK), nil

}

/*
GetWorkerPools gets worker pools

Get worker pools
*/
func (a *Client) GetWorkerPools(ctx context.Context, params *GetWorkerPoolsParams) (*GetWorkerPoolsOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkerPools",
		Method:             "GET",
		PathPattern:        "/jobservice/pools",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkerPoolsReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkerPoolsOK), nil

}

/*
GetWorkers gets workers

Get workers in current pool
*/
func (a *Client) GetWorkers(ctx context.Context, params *GetWorkersParams) (*GetWorkersOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getWorkers",
		Method:             "GET",
		PathPattern:        "/jobservice/pools/{pool_id}/workers",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetWorkersReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetWorkersOK), nil

}

/*
ListJobQueues lists job queues

list job queue
*/
func (a *Client) ListJobQueues(ctx context.Context, params *ListJobQueuesParams) (*ListJobQueuesOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listJobQueues",
		Method:             "GET",
		PathPattern:        "/jobservice/queues",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListJobQueuesReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListJobQueuesOK), nil

}

/*
StopRunningJob stops running job

Stop running job
*/
func (a *Client) StopRunningJob(ctx context.Context, params *StopRunningJobParams) (*StopRunningJobOK, error) {

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stopRunningJob",
		Method:             "PUT",
		PathPattern:        "/jobservice/jobs/{job_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StopRunningJobReader{formats: a.formats},
		AuthInfo:           a.authInfo,
		Context:            ctx,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StopRunningJobOK), nil

}