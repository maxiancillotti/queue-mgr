package datasum

import (
	"fmt"

	"queue-mgr/internal/jobs"
	"queue-mgr/internal/service"
)

func NewDataSumWorker() service.Worker {
	return &datasum{}
}

type datasum struct{}

func (w *datasum) Work(j *jobs.Job) {
	sum := 0

	for _, v := range j.Data {
		sum += v
	}

	j.SetResult(fmt.Sprintf("Processed successfully. Sum = %d", sum))
}
