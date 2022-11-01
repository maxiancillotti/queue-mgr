package mock

import (
	"context"

	"github.com/maxiancillotti/queue-mgr/internal/jobs"
)

type MockDispatcher struct{}

func (d *MockDispatcher) Start(ctx context.Context) {
}

func (d *MockDispatcher) Wait() {
}

func (d *MockDispatcher) Enqueue(job jobs.Job) {
}
