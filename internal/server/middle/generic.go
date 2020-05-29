package middle

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
)

// RegisterGenericHandlers appends generic error generators for unroutable
// requests to the global router.
func RegisterGenericHandlers(r *mux.Router) {
	r.MethodNotAllowedHandler = MetricAgg(
		RequestCtxProvider(midl.JSONAdapter(New405Handler())))
	r.NotFoundHandler = MetricAgg(
		RequestCtxProvider(midl.JSONAdapter(New404Handler())))
}
