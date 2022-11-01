package service

import (
	"context"

	"queue-mgr/internal/jobs"
)

type Dispatcher interface {
	// Starts a dispatcher that stops when it receives a value from ctx.Done.
	Start(ctx context.Context)

	// Use Wait to block until the dispatcher receives the signal to stop.
	Wait()

	// Enqueues a job.
	// Will block when the quantity of enqueued jobs has already reached to the maximum size,
	// until other job has been processed and frees up space in the queue.
	Enqueue(job jobs.Job)
}
