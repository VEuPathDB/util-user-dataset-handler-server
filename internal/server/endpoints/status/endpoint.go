package status

import (
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	// Std lib
	"net/http"

	// External
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	. "github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	tknKey  = "token"
	urlPath = "/process/dataset/{" + tknKey + "}/status"

	errInvalidState = "Invalid state: unknown data stored in cache for this " +
		"process"
)

func NewStatusEndpoint(o *config.Options, meta, upload *cache.Cache) types.Endpoint {
	return &statusEndpoint{
		opts:   o,
		meta:   meta,
		upload: upload,
	}
}

type statusEndpoint struct {
	log  *logrus.Entry
	opts *config.Options
	meta *cache.Cache
	upload *cache.Cache
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
	unkwn, ok := s.upload.Get(token)

	if !ok {
		return MakeResponse(http.StatusOK, job.StorableDetails{
			Token:  token,
			Status: job.StatusNotStarted,
		})
	}

	if det, ok := unkwn.(job.Details); ok {
		return MakeResponse(http.StatusOK, det.StorableDetails)
	}

	if det, ok := unkwn.(job.StorableDetails); ok {
		return MakeResponse(http.StatusOK, det)
	}

	GetCtxLogger(req).WithField("status", http.StatusInternalServerError).
		Error(errInvalidState)
	return svc.ServerError(errInvalidState)
}



