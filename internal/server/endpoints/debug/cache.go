package debug

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
	"github.com/gorilla/mux"
	"net/http"
)

const (
	urlCache = "/debug/cache"
)

func NewDebugEndpoint() types.Endpoint {
	return &debug{}
}

type debug struct{}

func (debug) Register(r *mux.Router) {
	r.Path(urlCache).
		Methods(http.MethodGet).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(midl.JSONAdapter(
			midl.MiddlewareFunc(HandleCache)))))
}

func HandleCache(midl.Request) midl.Response {
	out := make(map[string]map[string]interface{}, 3)

	tmp := cache.AllDetails()

	out["jobDetails"] = make(map[string]interface{}, len(tmp))
	for k, v := range tmp {
		out["jobDetails"][k] = v.Object
	}

	tmp = cache.AllMetadata()
	out["metadata"] = make(map[string]interface{}, len(tmp))
	for k, v := range tmp {
		out["metadata"][k] = v.Object
	}

	tmp = cache.AllHistoricalDetails()
	out["history"] = make(map[string]interface{}, len(tmp))
	for k, v := range tmp {
		out["history"][k] = v.Object
	}

	return midl.NewResponse().
		SetCode(http.StatusOK).
		SetBody(out)
}
