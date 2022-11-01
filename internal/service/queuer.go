package service

import "queue-mgr/internal/jobs"

// Queue memory repo
type Queuer interface {
	Enqueue(job *jobs.Job)
	Dequeue()

	RetrieveQueue() []*jobs.Job
	RetrievePendingQueue() []*jobs.Job
	RetrieveProcessedQueue() []*jobs.Job
}
