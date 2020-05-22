package middle

import (
	"strconv"
	"time"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/stats"
)

func NewTimer(next midl.Middleware) midl.MiddlewareFunc {
	return func(req midl.Request) midl.Response {
		start := time.Now()
		res   := next.Handle(req)
		recordTime(start, GetCtxLogger(req), res)
		return res
	}
}

func recordTime(start time.Time, logger *logrus.Entry, res midl.Response) {
	dur := time.Since(start)
	logger.WithField("duration", dur.String()).
		Info("Request completed with code " + strconv.Itoa(res.Code()))
	stats.GetServerStatus().RecordTime(dur)
}
