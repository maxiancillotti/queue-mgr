package middlewares

import (
	"net/http"

	"queue-mgr/internal/mock"
)

var (
	pstrMock mock.PresenterMock

	authTest = NewAuthController(&pstrMock)
)

func handlerMockforMDW(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusOK)
}

func handlerMockforMDWStatus500(rw http.ResponseWriter, req *http.Request) {
	rw.WriteHeader(http.StatusInternalServerError)
}
