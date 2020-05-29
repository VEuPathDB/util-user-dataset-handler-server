package api

import (
	"io"
	"net/http"
	"os"

	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
	"github.com/gorilla/mux"
)

// DocsEndpointPath defines the URL path to the API docs endpoint.
const DocsEndpointPath = "/api"

// NewAPIEndpoint returns a new Endpoint instance for the API docs endpoint.
func NewAPIEndpoint() types.Endpoint {
	return &docsEndpoint{}
}

type docsEndpoint struct{}

func (d *docsEndpoint) Register(r *mux.Router) {
	r.Path(DocsEndpointPath).
		Methods(http.MethodGet).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(http.HandlerFunc(
			func(out http.ResponseWriter, in *http.Request) { d.handle(out) }))))
}

func (d *docsEndpoint) handle(out http.ResponseWriter) {
	file, err := os.Open("static-content/api.html")
	if err != nil {
		out.WriteHeader(http.StatusInternalServerError)

		_, _ = out.Write([]byte(err.Error()))

		return
	}
	defer file.Close()

	out.Header().Set(xhttp.HeaderContentType, "text/html")
	out.WriteHeader(http.StatusOK)

	_, _ = io.Copy(out, file)
}
