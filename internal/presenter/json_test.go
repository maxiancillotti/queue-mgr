package presenter

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/maxiancillotti/queue-mgr/internal/handlers/presenter"
	"github.com/maxiancillotti/queue-mgr/internal/jobs"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPresenter presenter.Presenter = NewJSONPresenter()
)

//////////////////////////////////////////////

func TestErrorResp(t *testing.T) {

	// Initialization
	rwr := httptest.NewRecorder()

	errMsg := errors.New("ErrorMsg")

	// Execution
	testPresenter.PresentErrResponse(rwr, http.StatusBadRequest, errMsg)

	// Check
	statusCode := rwr.Result().StatusCode
	assert.Equal(t, http.StatusBadRequest, statusCode)

	var errResp ErrorResp
	err := json.NewDecoder(rwr.Result().Body).Decode(&errResp)
	assert.Nil(t, err)

	assert.Equal(t, errMsg.Error(), errResp.ErrMsg)
	assert.Equal(t, http.StatusBadRequest, errResp.HttpStatusCode)
}

func TestPresentResponse(t *testing.T) {

	t.Run("Success", func(t *testing.T) {
		// Initialization
		rwr := httptest.NewRecorder()

		jobResp := jobs.Job{Name: "testName"}

		// Execution
		testPresenter.PresentResponse(rwr, http.StatusOK, jobResp)

		// Check
		statusCode := rwr.Result().StatusCode
		assert.Equal(t, http.StatusOK, statusCode)

		var respBodyCheck jobs.Job
		err := json.NewDecoder(rwr.Result().Body).Decode(&respBodyCheck)
		assert.Nil(t, err)

		assert.Equal(t, jobResp.Name, respBodyCheck.Name)
	})

	t.Run("Panic", func(t *testing.T) {
		// Initialization
		rwr := httptest.NewRecorder()

		nonJsonVar := make(chan int) // var that doesn't support JSON encoding

		var panicMsgOut interface{}

		defer func() {
			if panicMsgOut = recover(); panicMsgOut != nil {
				t.Log("Panic recovered ok")

				// Check
				panicErr, ok := panicMsgOut.(error)
				assert.True(t, ok)

				assert.Contains(t, panicErr.Error(), errMsgPanicCannotEncodeResponseAsJSON)
			}
		}()

		// Execution
		testPresenter.PresentResponse(rwr, http.StatusOK, nonJsonVar)
	})
}
