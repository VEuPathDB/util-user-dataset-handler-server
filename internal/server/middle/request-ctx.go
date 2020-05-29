package middle

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"github.com/VEuPathDB/util-exporter-server/internal/service/rid"
)

func RequestCtxProvider() midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		if rid, err := rid.AssignRID(req); err != nil {
			log.Logger().WithField("endpoint", req.RawRequest().URL.Path).
				Error("Failed to generate request ID")
			return svc.ServerError("failed to generate request id")
		} else {
			logger.AddFields(rid, map[string]interface{}{
				"endpoint": req.RawRequest().URL.Path,
				"method":   req.RawRequest().Method,
			})
		}

		return nil
	}
}
