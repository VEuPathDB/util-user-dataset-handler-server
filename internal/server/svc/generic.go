package svc

import (
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
)

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
	Message *string         `json:"message,omitempty"`
}

func NotFound(msg string) midl.Response {
	return midl.MakeResponse(http.StatusNotFound, &SadResponse{
		Status:  StatusNotFound,
		Message: msg,
	})
}

func BadRequest(msg string) midl.Response {
	return midl.MakeResponse(http.StatusBadRequest, &SadResponse{
		Status:  StatusBadRequest,
		Message: msg,
	})
}

func ServerError(msg string) midl.Response {
	return midl.MakeResponse(http.StatusInternalServerError, &SadResponse{
		Status:  StatusServerErr,
		Message: msg,
	})
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
