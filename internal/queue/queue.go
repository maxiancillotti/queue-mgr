package queue

import "queue-mgr/internal/jobs"

type queue struct {
	jobs          []jobs.Job
	lastIdx       int
	lastProcessed int
}

func (q *queue) Enqueue(job jobs.Job) {

	q.jobs = append(q.jobs, job)
	q.lastIdx = len(q.jobs) - 1
	q.jobs[q.lastIdx].ID = q.lastIdx
}

func (q *queue) ProcessNext() {

}
