package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"queue-mgr/internal/handlers/internal"
	"queue-mgr/internal/handlers/presenter"
)

type AuthMDW interface {

	// Enables the queuing of jobs when Authorization header is set to "allow"
	AuthorizationAllow(next http.HandlerFunc) http.HandlerFunc
}

type authMDW struct {
	presenter presenter.Presenter
}

func NewAuthController(presenter presenter.Presenter) AuthMDW {
	return &authMDW{
		presenter: presenter,
	}
}

func (c *authMDW) AuthorizationAllow(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		expectedAuthorizationHeader := "allow"

		authHeader := strings.ToLower(req.Header.Get("Authorization"))
		if authHeader != expectedAuthorizationHeader {
			rw.Header().Add("WWW-Authenticate", expectedAuthorizationHeader)
			c.presenter.PresentErrResponse(rw, http.StatusUnauthorized, errors.New(internal.ErrMsgAuthorizationHeaderInvalid))
			return
		}

		next.ServeHTTP(rw, req)
	})
}
