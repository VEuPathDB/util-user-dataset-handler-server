package middle

import (
	"fmt"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
)

const (
	errBadContentType = "Unsupported content type '%s'.  Must be " +
		"'application/json'."
	errNoContentType = "Missing required Content-Type header."
)

// JSONContentFilter creates a middleware layer that enforces both the existence
// of the Content-Type header and that the header equals "application/json".
func JSONContentFilter() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		log := logger.ByRequest(req).WithField("status", http.StatusBadRequest)

		val, ok := req.Header("Content-Type")

		if !ok {
			log.Info(errNoContentType)

			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: errNoContentType,
			})
		}

		if val != "application/json" {
			msg := fmt.Sprintf(errBadContentType, val)

			log.Info(msg)

			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: msg,
			})
		}

		return nil
	}
}
