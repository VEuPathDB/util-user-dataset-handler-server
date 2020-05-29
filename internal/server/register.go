package server

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	// Std Lib
	"net/http"

	// External
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/api"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/health"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/options"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/status"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

type Server interface {
	RegisterEndpoints()
	Serve() error
}

func NewServer(o *config.Options) Server {
	return &server{mux.NewRouter(), o, log.Logger()}
}

type server struct {
	router  *mux.Router
	options *config.Options
	logger  *logrus.Entry
}

func (s *server) Serve() error {
	http.Handle("/", s.router)
	s.logger.Info("Server started.  Listening on port ", s.options.Port)
	return http.ListenAndServe(s.options.GetUsablePort(), nil)
}

func (s *server) RegisterEndpoints() {
	// Custom 404 & 405 handlers
	middle.RegisterGenericHandlers(s.router)

	s.router.Path("/metrics").
		Methods(http.MethodGet).
		Handler(promhttp.Handler())

	// Serve API docs
	api.NewApiEndpoint().Register(s.router)

	// Health Endpoint
	health.Register(s.router, s.options)

	// Options Endpoint
	options.Register(s.router, s.options)

	job.NewJobCreateEndpoint().Register(s.router)
	job.NewUploadEndpoint(s.options).Register(s.router)

	status.NewStatusEndpoint(s.options).Register(s.router)
}
