package xio

type ResponseStatus string

const (
	StatusOK         ResponseStatus = "ok"
	StatusNotFound   ResponseStatus = "not-found"
	StatusBadRequest ResponseStatus = "bad-request"
	StatusBadInput   ResponseStatus = "invalid-input"
	StatusServerErr  ResponseStatus = "server-error"
	StatusBadMethod  ResponseStatus = "bad-http-method"
)

type HappyResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message,omitempty"`
}

type SadResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
}

type ValidationResponse struct {
	Status  ResponseStatus      `json:"status"`
	Message string              `json:"message,omitempty"`
	Reasons map[string][]string `json:"reasons"`
}
