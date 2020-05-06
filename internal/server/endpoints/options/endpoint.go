package options

import (
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

const path = "/config"

func Register(r *mux.Router, o *config.Options) {
	r.Path(path).
		Methods(http.MethodGet).
		Handler(midl.JSONAdapter(middle.NewLogProvider(middle.NewTimer(
			func(log *logrus.Entry) midl.Middleware {
				return midl.MiddlewareFunc(func(request midl.Request) midl.Response {
					return midl.MakeResponse(http.StatusOK, o)
				})
			},
		))))
}
