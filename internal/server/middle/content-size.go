package middle

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
)

const (
	errBadLengthVal = "Content-Length value was not a valid unsigned integer."
	errNoLengthVal  = "Missing required header Content-Length."
	errTooBig       = "Content-Length %d is larger than the allowed size for this endpoint: %d"
)

// ContentLengthFilter constructs a middleware filter that enforces the incoming
// request has a Content-Length header and that the value is below the given
// threshold.
func ContentLengthFilter(bytes uint64)  midl.Middleware {
	return &contentLengthFilter{bytes: bytes}
}

type contentLengthFilter struct {
	bytes  uint64
}

func (c *contentLengthFilter) Handle(req midl.Request) midl.Response {
	log := logger.ByRequest(req).WithField("status", http.StatusBadRequest)

	if val, ok := req.Header(xhttp.HeaderContentLength); ok {
		size, err := strconv.ParseUint(val, 10, 64)

		if err != nil {
			log.Info(errBadLengthVal)
			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: errBadLengthVal,
			})
		}

		if size > c.bytes {
			msg := fmt.Sprintf(errTooBig, size, c.bytes)
			log.Info(msg)
			return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
				Status:  svc.StatusBadRequest,
				Message: msg,
			})
		}
	} else {
		log.Info(errNoLengthVal)
		return midl.MakeResponse(http.StatusBadRequest, &svc.SadResponse{
			Status:  svc.StatusBadRequest,
			Message: errNoLengthVal,
		})
	}

	return nil
}
