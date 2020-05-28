package api

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/xhttp"
	"github.com/gorilla/mux"
	"io"
	"net/http"
	"os"
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
		}))
}

