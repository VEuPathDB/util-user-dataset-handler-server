package svc

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

const (
	err405 = "Method not allowed."
	err404 = "Resource not found."
)

func New404Handler(logger *logrus.Entry) midl.Middleware {
	return midl.MiddlewareFunc(func(midl.Request) midl.Response {
		return midl.MakeResponse(http.StatusNotFound, &SadResponse{
			Status:  StatusNotFound,
			Message: err404,
		})
	})
}

func New405Handler(logger *logrus.Entry) midl.Middleware {
	return midl.MiddlewareFunc(func(midl.Request) midl.Response {
		return midl.MakeResponse(http.StatusMethodNotAllowed, &SadResponse{
			Status:  StatusBadMethod,
			Message: err405,
		})
	})
}
