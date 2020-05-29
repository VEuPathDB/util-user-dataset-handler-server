package middle

import (
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	errTknNotFound = "No process found for ID %s.  Either this process was " +
		"never started or has timed out."
)

// JobIdValidator is a middleware filter that rejects requests that have invalid
// jobId url parameters.
func JobIdValidator(
	tknKey string,
	next midl.Middleware,
) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		log := logger.ByRequest(req)
		log.Debug("Validating job id")
		token := mux.Vars(req.RawRequest())[tknKey]

		if !cache.HasMetadata(token) {
			errTxt := fmt.Sprintf(errTknNotFound, token)
			log.WithField("status", http.StatusNotFound).
				Info(errTxt)
			return svc.NotFound(errTxt)
		}

		return next.Handle(req)
	}
}
