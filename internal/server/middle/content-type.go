package middle

import (
	"fmt"
	"net/http"

	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
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
				return midl.MakeResponse(http.StatusBadRequest, &xio.SadResponse{
					Status:  xio.StatusBadRequest,
					Message: fmt.Sprintf(errBadContentType, val),
				})
			}
		} else {
			return midl.MakeResponse(http.StatusBadRequest, &xio.SadResponse{
				Status:  xio.StatusBadRequest,
				Message: errNoContentType,
			})
		}

		return nil
	}
}
