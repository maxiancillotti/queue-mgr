package service

import "github.com/maxiancillotti/queue-mgr/internal/jobs"

type Worker interface {
	Work(j *jobs.Job)
}
