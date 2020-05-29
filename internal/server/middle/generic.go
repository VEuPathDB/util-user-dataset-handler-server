package middle

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
)

func RegisterGenericHandlers(r *mux.Router) {
	r.MethodNotAllowedHandler = MetricAgg(
		RequestCtxProvider(midl.JSONAdapter(New405Handler())))
	r.NotFoundHandler = MetricAgg(
		RequestCtxProvider(midl.JSONAdapter(New404Handler())))
}

