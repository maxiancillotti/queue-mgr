package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/maxiancillotti/queue-mgr/internal/handlers/internal"
	"github.com/maxiancillotti/queue-mgr/internal/handlers/presenter"
	"github.com/maxiancillotti/queue-mgr/internal/jobs"
	"github.com/maxiancillotti/queue-mgr/internal/service"

	"github.com/pkg/errors"
)

type JobsController interface {
	// Creates a Job into the Queue.
	// Receives jobs.Job body:
	/*
		{
			"name": string,
			"data": []int
		}

	*/
	// API RESOURCE URL /jobs
	POST(rw http.ResponseWriter, req *http.Request)

	// Retrieves a collection of all jobs.
	// API RESOURCE URL /jobs
	GETCollection(rw http.ResponseWriter, req *http.Request)

	// Retrieves a collection of all jobs filtered by status.
	// Receives jobs.Job.status body:
	/*
		{
			"status": string
		}

	*/
	// API RESOURCE URL /jobs
	GETCollectionByStatus(rw http.ResponseWriter, req *http.Request)
}

type jobsController struct {
	dispatcher service.Dispatcher
	queuer     service.Queuer
	presenter  presenter.Presenter
}

func NewJobsController(dispatcher service.Dispatcher, queuer service.Queuer, presenter presenter.Presenter) JobsController {
	return &jobsController{
		dispatcher: dispatcher,
		queuer:     queuer,
		presenter:  presenter,
	}
}

func (c *jobsController) POST(rw http.ResponseWriter, req *http.Request) {

	var job jobs.Job

	err := json.NewDecoder(req.Body).Decode(&job)
	if err != nil {
		c.presenter.PresentErrResponse(rw, http.StatusBadRequest, errors.Wrap(err, internal.ErrMsgCannotDecodeJsonReqBody))
		return
	}

	err = job.ValidateInput()
	if err != nil {
		c.presenter.PresentErrResponse(rw, http.StatusBadRequest, errors.Wrap(err, internal.ErrMsgInvalidInputBody))
		return
	}

	c.dispatcher.Enqueue(job)

	rw.WriteHeader(http.StatusCreated)
}

func (c *jobsController) GETCollection(rw http.ResponseWriter, req *http.Request) {

	jobsQueue := c.queuer.RetrieveQueue()
	c.presenter.PresentResponse(rw, http.StatusOK, jobsQueue)
}

func (c *jobsController) GETCollectionByStatus(rw http.ResponseWriter, req *http.Request) {

	var job jobs.Job

	err := json.NewDecoder(req.Body).Decode(&job)
	if err != nil {
		c.presenter.PresentErrResponse(rw, http.StatusBadRequest, errors.Wrap(err, internal.ErrMsgCannotDecodeJsonReqBody))
		return
	}

	err = job.ValidateStatusFilter()
	if err != nil {
		c.presenter.PresentErrResponse(rw, http.StatusBadRequest, errors.Wrap(err, internal.ErrMsgInvalidInputBody))
		return
	}

	var jobs []*jobs.Job

	if job.IsProcessed() {
		jobs = c.queuer.RetrieveProcessedQueue()
	} else {
		jobs = c.queuer.RetrievePendingQueue()
	}

	c.presenter.PresentResponse(rw, http.StatusOK, jobs)
}
