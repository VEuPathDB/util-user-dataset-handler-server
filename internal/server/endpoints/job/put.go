package job

import (
	"github.com/vulpine-io/bites/v1/pkg/bites"
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
)

const (
	tokenKey = "job-id"
	urlPath  = "/job/{" + tokenKey + "}"
)

func NewJobCreateEndpoint() types.Endpoint {
	return &metadataEndpoint{}
}

type metadataEndpoint struct{}

func (m *metadataEndpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodPut).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(
			midl.JSONAdapter(
				middle.JSONContentFilter(),
				middle.ContentLengthFilter(bites.SizeMebibyte),
				NewMetadataValidator(),
				m))))
}

func (m *metadataEndpoint) Handle(req midl.Request) midl.Response {
	meta := req.AdditionalContext()[dataCtxKey].(*Metadata)

	cache.PutMetadata(meta.Token, meta.Metadata)

	return midl.MakeResponse(http.StatusNoContent, nil)
}
