package metadata

import (
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

const path = "/process/metadata"

func Register(r *mux.Router, _ *config.Options, c *cache.Cache) {
	r.Path(path).
		Methods(http.MethodPost).
		Handler(midl.JSONAdapter(
			middle.NewJsonContentFilter(),
			middle.NewContentLengthFilter(util.SizeMebibyte),
			middle.NewLogProvider(middle.NewTimer(
				func(log *logrus.Entry) midl.Middleware {
					return NewMetadataValidator(NewMetadataEndpoint(c))
				}))))
}

// NewMetadataEndpoint returns a constructor for the
// metadata POST endpoint that is intended to be wrapped
// by at least the ValidationWrapper middleware.
func NewMetadataEndpoint(c *cache.Cache) func(*Metadata) midl.Middleware {
	return func(meta *Metadata) midl.Middleware {
		return &metadataEndpoint{meta, c}
	}
}

type metadataEndpoint struct {
	meta *Metadata
	cche *cache.Cache
}

func (m *metadataEndpoint) Handle(midl.Request) midl.Response {
	if e := m.cche.Add(m.meta.Token, *m.meta, cache.DefaultExpiration); e != nil {
		return midl.MakeResponse(http.StatusInternalServerError, &svc.SadResponse{
			Status:  svc.StatusServerErr,
			Message: "Failed to write metadata to cache:" + e.Error(),
		})
	}
	return midl.MakeResponse(http.StatusOK, &svc.HappyResponse{
		Status:  svc.StatusOK,
	})
}
