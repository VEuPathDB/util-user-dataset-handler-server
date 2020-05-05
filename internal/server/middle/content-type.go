package middle

import (
	"fmt"
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
		if val, ok := req.Header("Content-Type"); ok {
			if val != "application/json" {
				return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
					Status:  svc.StatusBadRequest,
					Message: fmt.Sprintf(errBadContentType, val),
				})
			}
		} else {
			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: errNoContentType,
			})
		}

		return nil
	}
}
