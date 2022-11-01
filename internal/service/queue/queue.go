package queue

import (
	"sync"

	"github.com/maxiancillotti/queue-mgr/internal/jobs"
	"github.com/maxiancillotti/queue-mgr/internal/service"
)

func NewQueuer() service.Queuer {
	return &queue{
		pendingJobs:   make([]*jobs.Job, 0),
		processedJobs: make([]*jobs.Job, 0),
	}
}

type queue struct {
	pendingJobs   []*jobs.Job
	processedJobs []*jobs.Job
	priority      int
	wg            sync.WaitGroup
}

func (q *queue) Enqueue(job *jobs.Job) {

	q.wg.Wait()
	q.wg.Add(1)
	defer q.wg.Done()

	q.pendingJobs = append(q.pendingJobs, job)

	q.priority++
	//q.pendingJobs[len(q.pendingJobs)-1].ID = q.priority
	job.ID = q.priority
}

func (q *queue) Dequeue() {

	q.wg.Wait()
	q.wg.Add(1)
	defer q.wg.Done()

	q.processedJobs = append(q.processedJobs, q.pendingJobs[0])
	q.pendingJobs = q.pendingJobs[1:]
}

func (q *queue) RetrieveQueue() []*jobs.Job {
	return append(q.processedJobs, q.pendingJobs...)
}

func (q *queue) RetrievePendingQueue() []*jobs.Job {
	return q.pendingJobs
}

func (q *queue) RetrieveProcessedQueue() []*jobs.Job {
	return q.processedJobs
}
