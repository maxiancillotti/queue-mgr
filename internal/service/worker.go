package service

import "queue-mgr/internal/jobs"

type Worker interface {
	Work(j *jobs.Job)
}
