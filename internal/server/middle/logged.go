package middle

import (
	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

type LoggedMiddlewareFn func(logger *logrus.Entry) midl.Middleware

func NewLogProvider(next LoggedMiddlewareFn) midl.MiddlewareFunc {
	util.Logger().Trace("LoggedMiddlewareFn")
	return func(req midl.Request) midl.Response {
		id := uuid.New().String()
		log := util.Logger().
			WithField("request-id", id).
			WithField("endpoint", req.RawRequest().URL.Path)

		log.Trace("Prepared logger for request")
		return next(log).Handle(req)
	}
}