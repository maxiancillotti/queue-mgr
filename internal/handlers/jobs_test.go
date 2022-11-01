package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"queue-mgr/internal/handlers/internal"
	"queue-mgr/internal/jobs"
	"queue-mgr/internal/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	dispatcherMock = mock.MockDispatcher{}
	queuerMock     = mock.QueuerMock{}
	presenterMock  = mock.PresenterMock{}

	testJobsController = NewJobsController(&dispatcherMock, &queuerMock, &presenterMock)
)

func TestJobsPOST(t *testing.T) {

	type testCase struct {
		name           string
		reqBody        *jobs.Job
		expectedStatus int
		expectedErr    error
	}

	var testCases []testCase

	testCases = append(testCases, testCase{

		name: "Success",
		reqBody: &jobs.Job{
			Name: "JobName",
			Data: []int{1, 2, 3},
		},
		expectedStatus: http.StatusCreated,
		expectedErr:    nil,
	})

	testCases = append(testCases, testCase{

		name:           "Nil request body",
		reqBody:        nil,
		expectedStatus: http.StatusBadRequest,
		expectedErr:    errors.New(internal.ErrMsgCannotDecodeJsonReqBody),
	})

	testCases = append(testCases, testCase{

		name: "Empty Job Name",
		reqBody: &jobs.Job{
			Name: "",
			Data: []int{1, 2, 3},
		},
		expectedStatus: http.StatusBadRequest,
		expectedErr:    errors.New(internal.ErrMsgInvalidInputBody),
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			// Initialization
			var url string = "http://localhost:8001/api/jobs"

			buf := new(bytes.Buffer)

			if test.reqBody != nil {
				err := json.NewEncoder(buf).Encode(test.reqBody)
				assert.Nil(t, err)
			}

			req := httptest.NewRequest(http.MethodPost, url, buf)
			rwr := httptest.NewRecorder()

			// Execution
			testJobsController.POST(rwr, req)

			// Check
			statusCode := rwr.Result().StatusCode
			assert.Equal(t, test.expectedStatus, statusCode)

			if test.expectedStatus > 299 {
				// error case
				var errResp mock.ResponseMockError
				err := json.NewDecoder(rwr.Result().Body).Decode(&errResp)
				assert.Nil(t, err)

				assert.Contains(t, errResp.Msg, test.expectedErr.Error())

			}
		})
	}

}

func TestJobsGETCollection(t *testing.T) {

	type testCase struct {
		name             string
		expectedStatus   int
		expectedRespBody []jobs.Job
	}

	var testCases []testCase

	job1 := jobs.Job{
		Name: "JobName1",
		ID:   1,
		Data: []int{1, 2, 3},
	}
	job1.SetResult("Result OK. Sum = 6")

	job2 := jobs.Job{
		Name: "JobName2",
		ID:   2,
		Data: []int{3, 2, 2},
	}
	job2.SetResult("Result OK. Sum = 7")

	job3 := jobs.Job{
		Name: "JobName3",
		ID:   3,
		Data: []int{3, 2, 4},
	}
	job3.SetStatusPending()

	testCases = append(testCases, testCase{

		name:             "Success",
		expectedStatus:   http.StatusOK,
		expectedRespBody: []jobs.Job{job1, job2, job3},
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			// Initialization
			var url string = "http://localhost:8001/api/jobs"

			req := httptest.NewRequest(http.MethodGet, url, nil)
			rwr := httptest.NewRecorder()

			// Execution
			testJobsController.GETCollection(rwr, req)

			// Check
			statusCode := rwr.Result().StatusCode
			assert.Equal(t, test.expectedStatus, statusCode)

			var respJobs []jobs.Job
			err := json.NewDecoder(rwr.Result().Body).Decode(&respJobs)
			assert.Nil(t, err)

			assert.ElementsMatch(t, respJobs, test.expectedRespBody)
		})
	}
}

func TestJobsGETCollectionByStatus(t *testing.T) {

	type testCase struct {
		name             string
		reqBody          *jobs.Job
		expectedStatus   int
		expectedRespBody []jobs.Job
		expectedErrMsg   error
	}

	var testCases []testCase

	job3 := jobs.Job{
		Name: "JobName3",
		ID:   3,
		Data: []int{3, 2, 4},
	}
	job3.SetStatusPending()

	job4 := jobs.Job{
		Name: "JobName4",
		ID:   4,
		Data: []int{3, 2, 5},
	}
	job4.SetStatusPending()

	inputStatusPending := &jobs.Job{}
	inputStatusPending.SetStatusPending()

	testCases = append(testCases, testCase{

		name:             "Success. Status Pending.",
		reqBody:          inputStatusPending,
		expectedStatus:   http.StatusOK,
		expectedRespBody: []jobs.Job{job3, job4},
	})

	job1 := jobs.Job{
		Name: "JobName1",
		ID:   1,
		Data: []int{1, 2, 3},
	}
	job1.SetResult("Result OK. Sum = 6")

	job2 := jobs.Job{
		Name: "JobName2",
		ID:   2,
		Data: []int{3, 2, 2},
	}
	job2.SetResult("Result OK. Sum = 7")

	inputStatusProcessed := &jobs.Job{}
	inputStatusProcessed.SetStatusProcessed()

	testCases = append(testCases, testCase{

		name:             "Success. Status Processed.",
		reqBody:          inputStatusProcessed,
		expectedStatus:   http.StatusOK,
		expectedRespBody: []jobs.Job{job1, job2},
	})

	testCases = append(testCases, testCase{

		name:             "Request body decoding error",
		reqBody:          nil,
		expectedStatus:   http.StatusBadRequest,
		expectedRespBody: nil,
		expectedErrMsg:   errors.New(internal.ErrMsgCannotDecodeJsonReqBody),
	})

	testCases = append(testCases, testCase{

		name:             "Invalid input error",
		reqBody:          &jobs.Job{},
		expectedStatus:   http.StatusBadRequest,
		expectedRespBody: nil,
		expectedErrMsg:   errors.New(internal.ErrMsgInvalidInputBody),
	})

	for _, test := range testCases {

		t.Run(test.name, func(t *testing.T) {

			// Initialization
			var url string = "http://localhost:8001/api/jobs"

			buf := new(bytes.Buffer)

			if test.reqBody != nil {
				err := json.NewEncoder(buf).Encode(test.reqBody)
				assert.Nil(t, err)
			}

			req := httptest.NewRequest(http.MethodGet, url, buf)
			rwr := httptest.NewRecorder()

			// Execution
			testJobsController.GETCollectionByStatus(rwr, req)

			// Check
			statusCode := rwr.Result().StatusCode
			if assert.Equal(t, test.expectedStatus, statusCode) {

				if statusCode > 299 {
					var errorResp mock.ResponseMockError
					err := json.NewDecoder(rwr.Result().Body).Decode(&errorResp)
					assert.Nil(t, err)

					assert.Contains(t, errorResp.Msg, test.expectedErrMsg.Error())

				} else {
					var respJobs []jobs.Job
					err := json.NewDecoder(rwr.Result().Body).Decode(&respJobs)
					assert.Nil(t, err)

					assert.ElementsMatch(t, respJobs, test.expectedRespBody)
				}
			}
		})
	}
}
