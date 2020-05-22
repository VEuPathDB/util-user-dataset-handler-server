package middle

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	errBadContentType = "Unsupported content type '%s'.  Must be " +
		"'application/json'."
	errNoContentType = "Missing required Content-Type header."
)

// NewJsonContentFilter creates a middleware layer that
// enforces both the existence of the Content-Type header
// and that the header equals "application/json".
func NewJsonContentFilter() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		log := req.AdditionalContext()[KeyLogger].(*logrus.Entry).
			WithField("status", http.StatusBadRequest)
		if val, ok := req.Header("Content-Type"); ok {
			if val != "application/json" {
				msg := fmt.Sprintf(errBadContentType, val)
				log.Info(msg)
				return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
					Status:  svc.StatusBadRequest,
					Message: msg,
				})
			}
		} else {
			log.Info(errNoContentType)
			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: errNoContentType,
			})
		}

		return nil
	}
}
