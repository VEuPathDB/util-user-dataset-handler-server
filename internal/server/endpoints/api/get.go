package api

import (
	"github.com/VEuPathDB/util-exporter-server/internal/stats"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
	"github.com/gorilla/mux"
)

const DocsEndpointPath = "/api"

func NewApiEndpoint() types.Endpoint {
	return &docsEndpoint{}
}

type docsEndpoint struct {}

func (d *docsEndpoint) Register(r *mux.Router) {
	r.Path(DocsEndpointPath).
		Methods(http.MethodGet).
		Handler(http.HandlerFunc(func(out http.ResponseWriter, in *http.Request) {
			start := time.Now()
			d.handle(out)
			stats.GetServerStatus().RecordTime(time.Since(start))
		}))
}

func (d * docsEndpoint) handle(out http.ResponseWriter)  {
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

