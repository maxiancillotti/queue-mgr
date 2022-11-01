package middlewares

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/maxiancillotti/queue-mgr/internal/handlers/internal"
	"github.com/maxiancillotti/queue-mgr/internal/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testPanicRecoverMDW = NewPanicRecoverMiddleware(&pstrMock)
)

func handlerMockforMDWWithPanic(rw http.ResponseWriter, req *http.Request) {
	panic("this is a panic")
}

func TestPanicRecover(t *testing.T) {

	// Initialization
	req := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/jobs", nil)
	rwr := httptest.NewRecorder()

	// Execution
	testPanicRecoverMDW.PanicRecover(handlerMockforMDWWithPanic).ServeHTTP(rwr, req)

	// Check
	statusCode := rwr.Result().StatusCode
	assert.Equal(t, http.StatusInternalServerError, statusCode)

	var errResp mock.ResponseMockError
	err := json.NewDecoder(rwr.Result().Body).Decode(&errResp)
	assert.Nil(t, err)

	assert.Contains(t, errResp.Msg, internal.ErrMsgInternalServerUnexpected)
	assert.Contains(t, errResp.Msg, "this is a panic")
}
