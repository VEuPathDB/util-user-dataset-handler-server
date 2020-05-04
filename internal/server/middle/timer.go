package middle

import (
	"time"

	"gopkg.in/foxcapades/go-midl.v1/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/stats"
)

func NewTimer(next midl.Middleware) midl.MiddlewareFunc {
	return func(request midl.Request) midl.Response {
		start := time.Now()
		defer func() { stats.GetServerStatus().RecordTime(time.Since(start)) }()

		return next.Handle(request)
	}
}
