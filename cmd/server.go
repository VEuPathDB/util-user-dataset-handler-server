package main

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/health"
	"github.com/VEuPathDB/util-exporter-server/internal/server/metadata"
	"github.com/VEuPathDB/util-exporter-server/internal/server/options"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"

	// Std lib
	"net/http"
	"time"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

var version = "untagged dev build"

func main() {
	options := new(config.Options)
	options.Version = version
	config.ParseCli(options)
	config.ParseOptions(options)

	prepareLogger(options)

	r := mux.NewRouter()

	registerRoutes(r, options)

	http.Handle("/", r)
	util.Logger().Fatal(http.ListenAndServe(options.GetUsablePort(), nil))
}

func prepareLogger(opts *config.Options) {
	logrus.SetFormatter(new(logrus.TextFormatter))
	util.SetLogger(logrus.StandardLogger().
		WithField("service", opts.ServiceName))
}

func registerRoutes(r *mux.Router, o *config.Options) {
	statusCache := cache.New(72*time.Hour, time.Hour)
	uploadCache := cache.New(4*time.Hour, time.Hour)

	// Custom 404 & 405 handlers
	svc.RegisterGenericHandlers(r)

	// Serve API docs
	r.Get("/").Handler(http.FileServer(http.Dir("./static-content")))

	// Health Endpoint
	health.Register(r, o)

	// Options Endpoint
	options.Register(r, o)

	// Metadata recording endpoint
	metadata.Register(r, o, statusCache)

	r.Path("/process/dataset/{token}").
		Methods(http.MethodPost).
		Handler()

	r.Get(server.StatusEndpointPath).
		Handler(server.NewStatusEndpoint(statusCache))
}