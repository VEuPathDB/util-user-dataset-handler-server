package xio

import (
	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
	"net/http"
)

const (
	errNoRequestId  = "failed to assign a request id, middleware must be missing"
	errBadRequestId = "request id was not a valid uint32 value"
)

func ErrNoRequestId() midl.Response {
	return midl.MakeResponse(http.StatusInternalServerError, &SadResponse{
		Status:  StatusServerErr,
		Message: errNoRequestId,
	})
}

func ErrBadRequestId() midl.Response {
	return midl.MakeResponse(http.StatusInternalServerError, &SadResponse{
		Status:  StatusServerErr,
		Message: errBadRequestId,
	})
}
