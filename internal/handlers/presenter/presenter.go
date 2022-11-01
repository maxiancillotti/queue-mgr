package presenter

import (
	"net/http"
)

type Presenter interface {
	PresentResponse(rw http.ResponseWriter, status int, responseBody interface{})
	PresentErrResponse(rw http.ResponseWriter, status int, err error)
}
