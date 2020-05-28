package server

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/api"
	// Std Lib
	"net/http"
	"time"

	// External
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	cache2 "github.com/VEuPathDB/util-exporter-server/internal/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
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
	metaCache   := cache2.NewMeta(cache.New(72*time.Hour, time.Hour))
	uploadCache := cache2.NewUpload(cache.New(4*time.Hour, time.Hour))

	// Custom 404 & 405 handlers
	middle.RegisterGenericHandlers(s.router)

	// Serve API docs
	api.NewApiEndpoint().Register(s.router)

	// Health Endpoint
	health.Register(s.router, s.options)

	// Options Endpoint
	options.Register(s.router, s.options)

	job.NewJobCreateEndpoint(metaCache).Register(s.router)
	job.NewUploadEndpoint(s.options, metaCache, uploadCache).Register(s.router)

	status.NewStatusEndpoint(s.options, metaCache, uploadCache).Register(s.router)
}
