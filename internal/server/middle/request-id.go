package middle

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/teris-io/shortid"
)

func RequestIdProvider() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if code, err := shortid.Generate(); err != nil {
			log.Logger().WithField("endpoint", req.RawRequest().URL.Path).
				Error("Failed to generate request ID")
			return svc.ServerError("failed to generate request id")
		} else {
			req.AdditionalContext()[KeyRequestId] = code
		}

		return nil
	}
}
