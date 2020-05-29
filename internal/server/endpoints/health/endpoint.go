package health

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	// Std lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/stats"
)

const (
	path = "/health"
)

func Register(r *mux.Router, o *config.Options) {
	r.Path(path).
		Methods(http.MethodGet).
		Handler(midl.JSONAdapter(middle.NewTimer(
			middle.RequestCtxProvider(),
			&healthEndpoint{o.Version})))
}

type healthEndpoint struct {
	version string
}

func (h *healthEndpoint) Handle(midl.Request) midl.Response {
	return midl.MakeResponse(http.StatusOK, Data{
		Status:  "healthy", // TODO: define unhealthy status
		Version: h.version,
		Details: stats.GetServerStatus().ToPublic(),
	})
}
