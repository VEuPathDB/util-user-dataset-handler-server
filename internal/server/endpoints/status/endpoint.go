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
	"github.com/VEuPathDB/util-exporter-server/internal/process"
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
		Handler(JSONAdapter(NewLogProvider(NewTimer(
		NewTokenValidator(tknKey, s.meta, s.LogWrapper)))))
}

func (s *statusEndpoint) LogWrapper(log *logrus.Entry) Middleware {
	s.log = log
	return s
}

func (s *statusEndpoint) Handle(req Request) Response {
	token := mux.Vars(req.RawRequest())[tknKey]
	unkwn, ok := s.upload.Get(token)
	if !ok {
		return MakeResponse(http.StatusOK, process.StorableDetails{
			Token:    token,
			Status:   process.StatusNotStarted,
		})
	}
	if det, ok := unkwn.(process.Details); ok {
		return MakeResponse(http.StatusOK, det.StorableDetails)
	}
	if det, ok := unkwn.(process.StorableDetails); ok {
		return MakeResponse(http.StatusOK, det)
	}
	return svc.ServerError(errInvalidState)
}



