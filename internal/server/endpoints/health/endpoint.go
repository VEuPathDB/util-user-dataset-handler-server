package health

import (
	// Std lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/pkg/meta"
)

const (
	path = "/health"
)

// Register appends the "/health" endpoint to the given router.
func Register(r *mux.Router) {
	r.Path(path).
		Methods(http.MethodGet).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(
			midl.JSONAdapter(&healthEndpoint{}))))
}

type healthEndpoint struct {}

func (h *healthEndpoint) Handle(midl.Request) midl.Response {
	return midl.MakeResponse(http.StatusOK, Data{
		Status:  "healthy", // TODO: define unhealthy status
		Version: meta.GetBuildMeta().Version,
	})
}
