package middlewares

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	"github.com/maxiancillotti/queue-mgr/internal/handlers/internal"
	"github.com/maxiancillotti/queue-mgr/internal/mock"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthorizationXAdmin(t *testing.T) {

	expectedAuthorizationHeader := "allow"

	// ADD TEST CASES

	type testCase struct {
		name                 string
		headerValueInput     string
		expectedOutputStatus int
		expectedOutputErr    error
	}

	table := make([]testCase, 0)

	table = append(table, testCase{
		name:                 "success",
		headerValueInput:     "ALLOW",
		expectedOutputStatus: http.StatusOK,
		expectedOutputErr:    nil,
	})

	table = append(table, testCase{
		name:                 "success",
		headerValueInput:     "allow",
		expectedOutputStatus: http.StatusOK,
		expectedOutputErr:    nil,
	})

	table = append(table, testCase{
		name:                 "Error: authorization header is empty",
		headerValueInput:     "",
		expectedOutputStatus: http.StatusUnauthorized,
		expectedOutputErr:    errors.New(internal.ErrMsgAuthorizationHeaderInvalid),
	})

	//////////////////////////////////////////////
	// TESTING BEGINS

	for _, test := range table {

		t.Run(test.name, func(t *testing.T) {

			// Initialization
			req := httptest.NewRequest(http.MethodPost, "http://localhost:8001/api/jobs", nil)
			rwr := httptest.NewRecorder()

			req.Header.Set("Authorization", test.headerValueInput)

			// Execution
			authTest.AuthorizationAllow(handlerMockforMDW).ServeHTTP(rwr, req)

			// Check
			statusCode := rwr.Result().StatusCode
			assert.Equal(t, test.expectedOutputStatus, statusCode)

			if test.expectedOutputStatus > 299 {

				var errResp mock.ResponseMockError
				err := json.NewDecoder(rwr.Result().Body).Decode(&errResp)
				assert.Nil(t, err)

				assert.Contains(t, errResp.Msg, test.expectedOutputErr.Error())

				if test.expectedOutputStatus == http.StatusUnauthorized {
					headerAuthenticate := rwr.Result().Header.Get("WWW-Authenticate")
					assert.Equal(t, expectedAuthorizationHeader, headerAuthenticate)
				}
			}
		})
	}
}
