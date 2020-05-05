package main

import (
	"net/http"
	"time"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/server"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

var version = "untagged dev build"

func main() {
	statusCache := cache.New(72*time.Hour, time.Hour)
	uploadCache := cache.New(4*time.Hour, time.Hour)

	r := mux.NewRouter()

	// Custom 404 handler for uniform responses
	r.NotFoundHandler = midl.JSONAdapter(server.New404Handler())

	// Custom 405 handler for uniform responses
	r.MethodNotAllowedHandler = midl.JSONAdapter(server.New405Handler())

	// Serve API docs
	r.Get("/").Handler(http.FileServer(http.Dir("./static-content")))

	// Health Endpoint
	r.Get("/health").
		Handler(midl.JSONAdapter(server.NewHealthEndpoint(version)))

	r.Path("/process/metadata").
		Methods(http.MethodPost).
		Handler(midl.JSONAdapter(
			middle.NewJsonContentFilter(),
			middle.NewContentLengthFilter(util.SizeMebibyte),
			middle.NewTimer(
				server.NewMetadataWrapper(server.NewMetadataEndpoint(statusCache)))))

	r.Path("/process/dataset/{token}").
		Methods(http.MethodPost).
		Handler()

	r.Get("/status/{token}").Handler(server.NewStatusEndpoint(statusCache))

	http.Handle("/", r)
}
