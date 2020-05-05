package middle

import (
	"time"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/stats"
)

func NewTimer(next midl.Middleware) midl.MiddlewareFunc {
	return func(request midl.Request) midl.Response {
		start := time.Now()
		defer func() { stats.GetServerStatus().RecordTime(time.Since(start)) }()

		return next.Handle(request)
	}
}
