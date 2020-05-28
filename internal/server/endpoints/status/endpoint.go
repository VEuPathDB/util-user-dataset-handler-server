package status

import (
	"github.com/VEuPathDB/util-exporter-server/internal/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	// Std lib
	"net/http"

	// External
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	. "github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

const (
	tknKey  = "token"
	urlPath = "/job/{" + tknKey + "}/status"
)

func NewStatusEndpoint(
	opts *config.Options,
	meta *cache.Meta,
	upload *cache.Upload,
) types.Endpoint {
	return &statusEndpoint{
		opts:   opts,
		meta:   meta,
		upload: upload,
	}
}

type statusEndpoint struct {
	log    *logrus.Entry
	opts   *config.Options
	meta   *cache.Meta
	upload *cache.Upload
}

func (s *statusEndpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodGet).
		Handler(JSONAdapter(
			RequestIdProvider(),
			LogProvider(),
			NewTimer(NewTokenValidator(tknKey, s.meta, s))))
}

func (s *statusEndpoint) Handle(req Request) Response {
	token := mux.Vars(req.RawRequest())[tknKey]

	if det, ok := s.upload.GetStorable(token); ok {
		return MakeResponse(http.StatusOK, det)
	}

	if det, ok := s.upload.GetDetails(token); ok {
		return MakeResponse(http.StatusOK, det.StorableDetails)
	}

	meta, _ := s.meta.Get(token)
	return MakeResponse(http.StatusOK, job.StorableDetails{
		UserID:   meta.Owner,
		Token:    token,
		Status:   job.StatusNotStarted,
		Projects: meta.Projects,
	})
}
