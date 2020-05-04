package main

import (
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/server"
)

func main() {
	statusCache := cache.New(72*time.Hour, time.Hour)
	uploadCache := cache.New(4*time.Hour, time.Hour)
	r := mux.NewRouter()

	r.Path("/process/metadata").
		Methods(http.MethodPost)
	r.Path("/process/{token}").
		Methods(http.MethodPost)
	r.Get("/status/{token}").Handler(server.NewStatusEndpoint(statusCache))

	http.Handle("/", r)
}
