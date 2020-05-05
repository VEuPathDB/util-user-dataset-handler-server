package middle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
)

const (
	errBadLengthVal = "Content-Length value was not a valid unsigned integer."
	errNoLengthVal  = "Missing required header Content-Length."
	errTooBig       = "Content-Length %d is larger than the allowed size for this endpoint: %d"
)

// NewContentLengthFilter constructs a middleware filter
// that enforces the incoming request has a Content-Length
// header and that the value is below the given threshold.
func NewContentLengthFilter(bytes uint64) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if val, ok := req.Header(xhttp.HeaderContentLength); ok {
			size, err := strconv.ParseUint(val, 10, 64)
			if err != nil {
				return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
					Status:  svc.StatusBadRequest,
					Message: errBadLengthVal,
				})
			}
			if size > bytes {
				return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
					Status:  svc.StatusBadRequest,
					Message: fmt.Sprintf(errTooBig, size, bytes),
				})
			}
		} else {
			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: errNoLengthVal,
			})
		}

		return nil
	}
}
