package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
)

const urlPath = "/job/{job-id}"

func NewJobCreateEndpoint(c *cache.Cache) types.Endpoint {
	return &metadataEndpoint{metadataCache: c}
}

type metadataEndpoint struct {
	metadataCache *cache.Cache
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

	if e := m.metadataCache.Add(meta.Token, meta.Metadata, cache.DefaultExpiration); e != nil {
		return midl.MakeResponse(http.StatusInternalServerError, &svc.SadResponse{
			Status:  svc.StatusServerErr,
			Message: "Failed to write metadata to cache:" + e.Error(),
		})
	}
	return midl.MakeResponse(http.StatusOK, &svc.HappyResponse{
		Status: svc.StatusOK,
	})
}
