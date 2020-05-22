package middle

import (
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"net/http"
)

const (
	errBadToken = "Process ID must be a valid UUID v4 value."
	errTknNotFound = "No process found for ID %s.  Either this process was " +
		"never started or has timed out."
)

func NewTokenValidator(
	tknKey string,
	meta *cache.Cache,
	next Middleware,
) MiddlewareFunc {
	return func(req Request) Response {
		log := GetCtxLogger(req)
		log.Debug("Validating job id")
		token := mux.Vars(req.RawRequest())[tknKey]

		if _, err := uuid.Parse(token); err != nil {
			log.WithField("status", http.StatusBadRequest).
				Info(errBadToken)
			return svc.BadRequest(errBadToken)
		}

		if _, ok := meta.Get(token); !ok {
			log.WithField("status", http.StatusNotFound).
				Info(errTknNotFound)
			return svc.NotFound(errTknNotFound)
		}

		return next.Handle(req)
	}
}
