package middle

import (
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

const (
	errBadToken = "Process ID must be a valid UUID v4 value."
	errTknNotFound = "No process found for ID %s.  Either this process was " +
		"never started or has timed out."
)

func NewTokenValidator(
	tknKey string,
	meta *cache.Cache,
	next LoggedMiddlewareFn,
) LoggedMiddlewareFn {
	return func(log *logrus.Entry) Middleware {
		return MiddlewareFunc(func(req Request) Response {
			token := mux.Vars(req.RawRequest())[tknKey]
			if _, err := uuid.Parse(token); err != nil {
				return svc.BadRequest(errBadToken)
			}

			if _, ok := meta.Get(token); !ok {
				return svc.NotFound(errTknNotFound)
			}

			return next(log).Handle(req)
		})
	}
}
