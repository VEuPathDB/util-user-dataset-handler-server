package server

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
	"net/http"
)

const (
	err405 = "Method not allowed."
	err404 = "Resource not found."
)

func New404Handler() midl.MiddlewareFunc {
	return func(midl.Request) midl.Response {
		return midl.MakeResponse(http.StatusNotFound, &xio.SadResponse{
			Status:  xio.StatusNotFound,
			Message: err404,
		})
	}
}

func New405Handler() midl.MiddlewareFunc {
	return func(midl.Request) midl.Response {
		return midl.MakeResponse(http.StatusMethodNotAllowed, &xio.SadResponse{
			Status:  xio.StatusBadMethod,
			Message: err405,
		})
	}
}
