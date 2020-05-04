package middle

import (
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"
	"net/http"
	"strconv"
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
				return midl.MakeResponse(http.StatusBadRequest, &xio.SadResponse{
					Status:  xio.StatusBadRequest,
					Message: errBadLengthVal,
				})
			}
			if size > bytes {
				return midl.MakeResponse(http.StatusBadRequest, &xio.SadResponse{
					Status:  xio.StatusBadRequest,
					Message: fmt.Sprintf(errTooBig, size, bytes),
				})
			}
		} else {
			return midl.MakeResponse(http.StatusBadRequest, &xio.SadResponse{
				Status:  xio.StatusBadRequest,
				Message: errNoLengthVal,
			})
		}

		return nil
	}
}
