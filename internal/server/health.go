package server

import (
	"net/http"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

	"github.com/VEuPathDB/util-exporter-server/internal/server/xio"
	"github.com/VEuPathDB/util-exporter-server/internal/stats"
)

// NewHealthEndpoint constructs a new /health endpoint
// handler.
func NewHealthEndpoint(version string) midl.Middleware {
	return &healthEndpoint{version: version}
}

type healthEndpoint struct {
	version string
}

func (h *healthEndpoint) Handle(midl.Request) midl.Response {
	return midl.MakeResponse(http.StatusOK, xio.HealthData{
		Status:  "healthy", // TODO: define unhealthy status
		Version: h.version,
		Details: stats.GetServerStatus().ToPublic(),
	})
}
