package mock

import (
	"encoding/json"
	"net/http"
)

type PresenterMock struct{}

type ResponseMockSuccess struct {
	Msg string `json:"response"`
}

type ResponseMockError struct {
	Msg string `json:"response"`
}

func (p *PresenterMock) PresentErrResponse(rw http.ResponseWriter, status int, errorResp error) {

	responseBody := ResponseMockError{errorResp.Error()}
	p.PresentResponse(rw, status, responseBody)
}

func (p *PresenterMock) PresentResponse(rw http.ResponseWriter, status int, responseBody interface{}) {

	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	json.NewEncoder(rw).Encode(responseBody)
}
