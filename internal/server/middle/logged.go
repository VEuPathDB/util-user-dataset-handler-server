package middle

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/sirupsen/logrus"
)

func LogProvider() midl.MiddlewareFunc {
	log.Logger().Trace("LoggedMiddlewareFn")
	return func(req midl.Request) midl.Response {
		log := log.Logger().
			WithField("request-id", req.AdditionalContext()["request-id"]).
			WithField("endpoint", req.RawRequest().URL.Path).
			WithField("method", req.RawRequest().Method)

		log.Debug("Prepared logger for request")
		req.AdditionalContext()[KeyLogger] = log

		return nil
	}
}

func GetCtxLogger(req midl.Request) *logrus.Entry {
	return req.AdditionalContext()[KeyLogger].(*logrus.Entry)
}