package svc

import (
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
)

const (
	err405 = "Method not allowed."
	err404 = "Resource not found."
)

func New404Handler() midl.Middleware {
	return midl.MiddlewareFunc(func(r midl.Request) midl.Response {
		logger.ByRequest(r).
			WithField("status", http.StatusNotFound).
			Info("Not found")
		return midl.MakeResponse(http.StatusNotFound, &SadResponse{
			Status:  StatusNotFound,
			Message: err404,
		})
	})
}

func New405Handler() midl.Middleware {
	return midl.MiddlewareFunc(func(r midl.Request) midl.Response {
		logger.ByRequest(r).
			WithField("status", http.StatusMethodNotAllowed).
			Info("Method not allowed")
		return midl.MakeResponse(http.StatusMethodNotAllowed, &SadResponse{
			Status:  StatusBadMethod,
			Message: err405,
		})
	})
}
