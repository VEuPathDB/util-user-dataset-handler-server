package server

import (
	// Std Lib
	"net/http"

	// External
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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

// Server defines a wrapper around the HTTP server functionality, specifically
// involving setting up and starting an HTTP server.
type Server interface {

	// RegisterEndpoints configures all routes for the server.
	RegisterEndpoints()

	// Serve starts the server instance.
	Serve() error
}

// NewServer returns a new instance of the Server type, configured with the
// given options.
func NewServer(cli config.CLIOptions, file config.FileOptions) Server {
	return &server{mux.NewRouter(), cli, file, log.Logger()}
}

type server struct {
	router      *mux.Router
	cliOptions  config.CLIOptions
	fileOptions config.FileOptions
	logger      *logrus.Entry
}

func (s *server) Serve() error {
	http.Handle("/", s.router)
	s.logger.Info("Server started.  Listening on port ", s.cliOptions.Port())

	return http.ListenAndServe(s.cliOptions.GetUsablePort(), nil)
}

func (s *server) RegisterEndpoints() {
	// Custom 404 & 405 handlers
	middle.RegisterGenericHandlers(s.router)

	s.router.Path("/metrics").
		Methods(http.MethodGet).
		Handler(promhttp.Handler())

	// Serve API docs
	api.NewAPIEndpoint().Register(s.router)

	// Health Endpoint
	health.Register(s.router)

	// Options Endpoint
	options.Register(s.router, s.cliOptions, s.fileOptions)

	job.NewJobCreateEndpoint().Register(s.router)
	job.NewUploadEndpoint(s.fileOptions).Register(s.router)
	status.NewStatusEndpoint(s.fileOptions).Register(s.router)
}
