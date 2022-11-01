package dispatcher

import (
	"context"
	"sync"
	"time"

	"queue-mgr/internal/jobs"
	"queue-mgr/internal/service"
)

// Dispatcher represents a job dispatcher.
type dispatcher struct {
	// Restrict the number of goroutines using buffered channel (as counting semaphor)
	semaphore chan struct{}

	// Buffer for maximum size of the queue
	jobBuffer chan *jobs.Job

	// Delay time between the process of two jobs
	timeBtwJobProc time.Duration

	// Process the job
	worker service.Worker

	// Needed to run background dequeuing process
	wg sync.WaitGroup
}

func (d *dispatcher) StartDequeue(ctx context.Context) {
	d.wg.Add(1)
	go d.loop(ctx)
}

func (d *dispatcher) Wait() {
	d.wg.Wait()
}

func (d *dispatcher) Enqueue(job *jobs.Job) {
	d.jobBuffer <- job
}

func (d *dispatcher) stop() {
	d.wg.Done()
}

func (d *dispatcher) loop(ctx context.Context) {
	var wg sync.WaitGroup
Loop:
	for {
		select {

		case <-ctx.Done():
			// Block until all the jobs finishes
			wg.Wait()
			break Loop

		case job := <-d.jobBuffer:
			// Increment the waitgroup
			wg.Add(1)

			// Decrement a semaphore count
			d.semaphore <- struct{}{}

			// Fire new worker goroutine
			go func(job *jobs.Job) {
				defer wg.Done()
				// After the job finished, increment a semaphore count
				defer func() { <-d.semaphore }()
				d.worker.Work(job)
			}(job)

			// Blocking state until time has passed
			<-time.After(d.timeBtwJobProc)
		}
	}
	d.stop()
}
