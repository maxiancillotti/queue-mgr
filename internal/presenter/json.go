package presenter

import (
	"encoding/json"
	"net/http"

	"queue-mgr/internal/handlers/presenter"

	"github.com/pkg/errors"
)

const (
	contentTypeHeaderKey       = "Content-Type"
	contentTypeHeaderValueJSON = "application/json"
)

type jsonPresenter struct{}

func NewJSONPresenter() presenter.Presenter {
	return &jsonPresenter{}
}

func (p *jsonPresenter) PresentErrResponse(rw http.ResponseWriter, status int, errResp error) {

	responseBody := ErrorResp{ErrMsg: errResp.Error(), HttpStatusCode: status}
	p.PresentResponse(rw, status, responseBody)
}

func (p *jsonPresenter) PresentResponse(rw http.ResponseWriter, status int, responseBody interface{}) {

	// Header first because Write method without calling WriteHeader
	// will write status 200 automatically.
	rw.Header().Set(contentTypeHeaderKey, contentTypeHeaderValueJSON)
	rw.WriteHeader(status)

	err := json.NewEncoder(rw).Encode(responseBody)
	if err != nil {
		panic(errors.Wrap(err, errMsgPanicCannotEncodeResponseAsJSON))
	}
}
