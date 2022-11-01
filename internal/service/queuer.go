package service

import "queue-mgr/internal/jobs"

type Queuer interface {
	Enqueue(job *jobs.Job)
	Dequeue()
}
