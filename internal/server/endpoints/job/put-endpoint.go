package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
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

func NewJobCreateEndpoint() types.Endpoint {
	return &metadataEndpoint{}
}

type metadataEndpoint struct {}

func (m *metadataEndpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodPut).
		Handler(midl.JSONAdapter(
			middle.RequestCtxProvider(),
			middle.NewTimer(
				middle.NewJsonContentFilter(),
				middle.NewContentLengthFilter(util.SizeMebibyte),
				NewMetadataValidator(),
				m,
			)))
}

func (m *metadataEndpoint) Handle(req midl.Request) midl.Response {
	meta := req.AdditionalContext()[dataCtxKey].(*Metadata)

	cache.PutMetadata(meta.Token, meta.Metadata)

	return midl.MakeResponse(http.StatusOK, &svc.HappyResponse{
		Status: svc.StatusOK,
	})
}
