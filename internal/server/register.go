package server

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"github.com/sirupsen/logrus"
	// Std Lib
	"net/http"
	"time"

	// External
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
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
	r *mux.Router
	o *config.Options
	l *logrus.Entry
}

func (s *server) Serve() error {
	http.Handle("/", s.r)
	return http.ListenAndServe(s.o.GetUsablePort(), nil)
}

func (s *server) RegisterEndpoints() {
	metaCache   := cache.New(72*time.Hour, time.Hour)
	uploadCache := cache.New(4*time.Hour, time.Hour)

	// Custom 404 & 405 handlers
	middle.RegisterGenericHandlers(s.r)

	// Serve API docs
	s.r.Path("/api").
		Methods(http.MethodGet).
		Handler(http.FileServer(http.Dir("./static-content")))

	// Health Endpoint
	health.Register(s.r, s.o)

	// Options Endpoint
	options.Register(s.r, s.o)

	job.NewJobCreateEndpoint(metaCache).Register(s.r)
	job.NewUploadEndpoint(s.o, metaCache, uploadCache).Register(s.r)

	status.NewStatusEndpoint(s.o, metaCache, uploadCache).Register(s.r)
}
