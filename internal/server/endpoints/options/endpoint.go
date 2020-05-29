package options

import (
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

const path = "/config"

func Register(r *mux.Router, o *config.Options) {
	r.Path(path).
		Methods(http.MethodGet).
		Handler(midl.JSONAdapter(
			middle.RequestCtxProvider(),
			middle.NewTimer(midl.MiddlewareFunc(func(request midl.Request) midl.Response {
					return midl.MakeResponse(http.StatusOK, o)
			}))))
}
