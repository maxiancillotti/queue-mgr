package presenter

const (
	errMsgPanicCannotEncodeResponseAsJSON = "cannot encode response body as a JSON"
)

type ErrorResp struct {
	ErrMsg         string `json:"error"`
	HttpStatusCode int    `json:"status"`
}
