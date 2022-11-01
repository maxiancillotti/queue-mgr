package jobs

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetResult(t *testing.T) {

	testJob := Job{
		Name:   "JobName",
		ID:     10,
		Status: statusPending,
		Data:   []int{1, 2, 3},
	}

	result := "result ok"
	testJob.SetResult(result)

	assert.Equal(t, result, string(testJob.Result))
	assert.Equal(t, statusProcessed, testJob.Status)
}

func TestSetStatusPending(t *testing.T) {

	testJob := Job{
		Name: "JobName",
		Data: []int{1, 2, 3},
	}
	testJob.SetStatusPending()

	assert.Equal(t, statusPending, testJob.Status)
}

func TestSetStatusProcessed(t *testing.T) {

	testJob := Job{
		Name: "JobName",
		Data: []int{1, 2, 3},
	}
	testJob.SetStatusProcessed()

	assert.Equal(t, statusProcessed, testJob.Status)
}

func TestValidateInput(t *testing.T) {

	type testCase struct {
		name                 string
		inputJob             Job
		expectedOutputErrMsg string
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		name: "Empty name",
		inputJob: Job{
			Name: "",
			Data: []int{1, 2, 3},
		},
		expectedOutputErrMsg: "name cannot be empty",
	})

	testCases = append(testCases, testCase{
		name: "Empty data",
		inputJob: Job{
			Name: "job name",
			Data: []int{},
		},
		expectedOutputErrMsg: "data cannot be empty",
	})

	testCases = append(testCases, testCase{
		name: "Nil data",
		inputJob: Job{
			Name: "job name",
			Data: nil,
		},
		expectedOutputErrMsg: "data cannot be empty",
	})

	testCases = append(testCases, testCase{
		name: "Success",
		inputJob: Job{
			Name: "job name",
			Data: []int{1, 2, 3},
		},
		expectedOutputErrMsg: "",
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			err := test.inputJob.ValidateInput()

			if test.expectedOutputErrMsg != "" {
				assert.Equal(t, test.expectedOutputErrMsg, err.Error())
			} else {
				assert.Nil(t, err)
			}
		})
	}
}

func TestValidateStatusFilter(t *testing.T) {

	type testCase struct {
		name              string
		inputJob          Job
		expectedOutputErr error
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		name: "Status Invalid",
		inputJob: Job{
			Status: jobStatus("invalidStatus"),
		},
		expectedOutputErr: errors.New("status is invalid"),
	})

	testCases = append(testCases, testCase{
		name: "Status OK",
		inputJob: Job{
			Status: statusPending,
		},
		expectedOutputErr: nil,
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			err := test.inputJob.ValidateStatusFilter()

			assert.Equal(t, test.expectedOutputErr, err)
		})
	}
}

func TestIsProcessed(t *testing.T) {

	type testCase struct {
		name           string
		inputJob       Job
		expectedOutput bool
	}

	var testCases []testCase

	testCases = append(testCases, testCase{
		name: "Is Processed",
		inputJob: Job{
			Status: statusProcessed,
		},
		expectedOutput: true,
	})

	testCases = append(testCases, testCase{
		name: "Is Not Processed",
		inputJob: Job{
			Status: statusPending,
		},
		expectedOutput: false,
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {
			isProcessed := test.inputJob.IsProcessed()

			assert.Equal(t, test.expectedOutput, isProcessed)
		})
	}
}
