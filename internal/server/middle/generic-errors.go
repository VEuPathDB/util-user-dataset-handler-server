package middle

import (
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
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

		promTotalRequests.WithLabelValues(
			r.RawRequest().URL.Path,
			r.RawRequest().Method,
			"404",
		).Inc()

		return midl.MakeResponse(http.StatusNotFound, &svc.SadResponse{
			Status:  svc.StatusNotFound,
			Message: err404,
		})
	})
}

func New405Handler() midl.Middleware {
	return midl.MiddlewareFunc(func(r midl.Request) midl.Response {
		logger.ByRequest(r).
			WithField("status", http.StatusMethodNotAllowed).
			Info("Method not allowed")

		promTotalRequests.WithLabelValues(
			r.RawRequest().URL.Path,
			r.RawRequest().Method,
			"405",
		).Inc()

		return midl.MakeResponse(http.StatusMethodNotAllowed, &svc.SadResponse{
			Status:  svc.StatusBadMethod,
			Message: err405,
		})
	})
}
