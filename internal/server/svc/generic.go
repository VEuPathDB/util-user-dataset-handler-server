package svc

import (
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
)

// ResponseStatus defines a service response status type.
type ResponseStatus string

// All possible response status types.
const (
	StatusNotFound   ResponseStatus = "not-found"
	StatusBadRequest ResponseStatus = "bad-request"
	StatusBadInput   ResponseStatus = "invalid-input"
	StatusServerErr  ResponseStatus = "server-error"
	StatusBadMethod  ResponseStatus = "bad-http-method"
)

// NotFound constructs a simple 404 response body.
func NotFound(msg string) midl.Response {
	return midl.MakeResponse(http.StatusNotFound, &SadResponse{
		Status:  StatusNotFound,
		Message: msg,
	}).SetHeader("Content-Type", "application/json")
}

// InvalidRequest constructs a simple 422 response body.
func InvalidRequest(msg string) midl.Response {
	return midl.MakeResponse(http.StatusUnprocessableEntity, &ValidationResponse{
		Status: StatusBadInput,
		Errors: ValidationBundle{
			General: []string{msg},
		},
	}).SetHeader("Content-Type", "application/json")
}

// BadRequest constructs a simple 400 response body.
func BadRequest(msg string) midl.Response {
	return midl.MakeResponse(http.StatusBadRequest, &SadResponse{
		Status:  StatusBadRequest,
		Message: msg,
	}).SetHeader("Content-Type", "application/json")
}

// ServerError constructs a simple 500 response body.
func ServerError(msg string) midl.Response {
	return midl.MakeResponse(http.StatusInternalServerError, &SadResponse{
		Status:  StatusServerErr,
		Message: msg,
	}).SetHeader("Content-Type", "application/json")
}

// SadResponse defines a generic error response body.
type SadResponse struct {
	Status  ResponseStatus `json:"status"`
	Message string         `json:"message"`
}

// ValidationResponse defines a 422 error response body.
type ValidationResponse struct {
	Status  ResponseStatus   `json:"status"`
	Errors  ValidationBundle `json:"reasons"`
}

type ValidationBundle struct {
	General []string            `json:"general"`
	ByKey   map[string][]string `json:"byKey"`
}
