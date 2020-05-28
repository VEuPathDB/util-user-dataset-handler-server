package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/gorilla/mux"
)
const (
	tokenKey = "job-id"
	urlPath = "/job/{" + tokenKey + "}"
)

func NewJobCreateEndpoint(c *cache.Meta) types.Endpoint {
	return &metadataEndpoint{meta: c}
}

type metadataEndpoint struct {
	meta *cache.Meta
}

func (m *metadataEndpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodPut).
		Handler(midl.JSONAdapter(
			middle.RequestIdProvider(),
			middle.LogProvider(),
			middle.NewJsonContentFilter(),
			middle.NewContentLengthFilter(util.SizeMebibyte),
			NewMetadataValidator(),
			middle.NewTimer(m)))
}

func (m *metadataEndpoint) Handle(req midl.Request) midl.Response {
	meta := req.AdditionalContext()["data"].(*Metadata)

	m.meta.Set(meta.Token, meta.Metadata)

	return midl.MakeResponse(http.StatusOK, &svc.HappyResponse{
		Status: svc.StatusOK,
	})
}
