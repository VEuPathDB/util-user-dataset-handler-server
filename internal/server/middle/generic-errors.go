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

// New404Handler constructs a generic handler for requests that have no matching
// registered route.
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

// New405Handler constructs a generic handler for requests that have no matching
// route for the request method (e.g. no handler for POST request to /api).
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
