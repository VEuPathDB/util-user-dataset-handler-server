package middle

import (
	"fmt"
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	errTknNotFound = "No process found for ID %s.  Either this process was " +
		"never started or has timed out."
)

func NewTokenValidator(
	tknKey string,
	meta *cache.Meta,
	next Middleware,
) MiddlewareFunc {
	return func(req Request) Response {
		log := GetCtxLogger(req)
		log.Debug("Validating job id")
		token := mux.Vars(req.RawRequest())[tknKey]

		if _, ok := meta.Get(token); !ok {
			errTxt := fmt.Sprintf(errTknNotFound, token)
			log.WithField("status", http.StatusNotFound).
				Info(errTxt)
			return svc.NotFound(errTxt)
		}

		return next.Handle(req)
	}
}
