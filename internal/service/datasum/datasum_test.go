package datasum

import (
	"queue-mgr/internal/jobs"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testWorker = NewDataSumWorker()
)

func TestWork(t *testing.T) {

	testJob := jobs.Job{
		Data: []int{1, 2, 3},
	}
	testResult := "Processed successfully. Sum = 6"

	testWorker.Work(&testJob)

	assert.Equal(t, testResult, string(testJob.Result))
}
