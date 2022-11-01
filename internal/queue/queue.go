package queue

import (
	"queue-mgr/internal/jobs"
)

type queue struct {
	jobs          []jobs.Job
	lastProcessed int
}

func (q *queue) Enqueue(job jobs.Job) {

	q.jobs = append(q.jobs, job)
	queueLen := len(q.jobs)
	q.jobs[queueLen-1].ID = queueLen
}

func (q *queue) RetrieveQueue() []jobs.Job {
	return q.jobs
}

func (q *queue) RetrievePendingQueue() []jobs.Job {
	return q.jobs[q.lastProcessed:]
}

func (q *queue) RetrieveProcessedQueue() []jobs.Job {
	return q.jobs[:q.lastProcessed+1]
}

func (q *queue) RetrieveNextToProcess() jobs.Job {
	// OJO
	// time.After + select + inyectar parámetro

	return q.jobs[q.lastProcessed+1]
}
