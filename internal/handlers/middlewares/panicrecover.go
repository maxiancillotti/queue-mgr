package middlewares

import (
	"fmt"
	"net/http"

	"github.com/maxiancillotti/queue-mgr/internal/handlers/internal"
	"github.com/maxiancillotti/queue-mgr/internal/handlers/presenter"

	"github.com/pkg/errors"
)

type PanicRecoverMiddleware interface {
	PanicRecover(next http.HandlerFunc) http.HandlerFunc
}

func NewPanicRecoverMiddleware(presenter presenter.Presenter) PanicRecoverMiddleware {
	return &panicRecoverMiddleware{presenter: presenter}
}

type panicRecoverMiddleware struct {
	presenter presenter.Presenter
}

func (m *panicRecoverMiddleware) PanicRecover(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {

		defer func() {
			if pncErr := recover(); pncErr != nil {
				err := errors.New(fmt.Sprint(pncErr))

				// this should also log the errors with zap
				m.presenter.PresentErrResponse(rw, http.StatusInternalServerError, errors.Wrap(err, internal.ErrMsgInternalServerUnexpected))
			}
		}()

		next.ServeHTTP(rw, req)
	})
}
