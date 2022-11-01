package dispatcher

import (
	"queue-mgr/internal/jobs"
	"queue-mgr/internal/service"

	"runtime"
	"time"
)

const (
	defaultQueueLenght    = 1000
	defaultTimeBtwJobProc = 10 * time.Second
)

var (
	defaultMaxWorkers = runtime.NumCPU()
)

type dispatcherBuilder struct {
	maxWorkers     int
	queueLen       int
	timeBtwJobProc time.Duration
}

func NewDispatcherBuilder() *dispatcherBuilder {
	return &dispatcherBuilder{
		maxWorkers:     defaultMaxWorkers,
		queueLen:       defaultQueueLenght,
		timeBtwJobProc: defaultTimeBtwJobProc,
	}
}

// Set maximum number of goroutines that can work concurrently. Default: number of logical CPUs usable.
func (b *dispatcherBuilder) SetMaxWorkers(mw int) *dispatcherBuilder {
	b.maxWorkers = mw
	return b
}

// Set maximum size of the queue. Default: 1000.
func (b *dispatcherBuilder) SetQueueLen(ql int) *dispatcherBuilder {
	b.queueLen = ql
	return b
}

// Set delay time between the process of two jobs. Default: 10 seconds.
func (b *dispatcherBuilder) SetTimeBetweenJobProcesses(timebtw time.Duration) *dispatcherBuilder {
	b.timeBtwJobProc = timebtw
	return b
}

// Builds new instance of job dispatcher with the config set or its defaults.
// Needs a worker to process the job and a queuer to save the queue.
func (b *dispatcherBuilder) BuildDispatcher(worker service.Worker, queuer service.Queuer) service.Dispatcher {
	return &dispatcher{
		semaphore:      make(chan struct{}, b.maxWorkers),
		jobBuffer:      make(chan *jobs.Job, b.queueLen),
		timeBtwJobProc: b.timeBtwJobProc,
		worker:         worker,
		queuer:         queuer,
	}
}
