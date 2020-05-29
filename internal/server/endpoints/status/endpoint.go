package status

import (
	// Std lib
	"net/http"

	// External
	. "github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
)

const (
	tknKey  = "token"
	urlPath = "/job/{" + tknKey + "}/status"
)

func NewStatusEndpoint(opts *config.Options) types.Endpoint {
	return &statusEndpoint{opts: opts}
}

type statusEndpoint struct {
	opts *config.Options
}

func (s *statusEndpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodGet).
		Handler(JSONAdapter(
			middle.RequestCtxProvider(),
			middle.NewTimer(s)))
}

func (s *statusEndpoint) Handle(req Request) Response {
	jobId := mux.Vars(req.RawRequest())[tknKey]

	// Is the job in progress
	if det, ok := cache.GetDetails(jobId); ok {
		MakeResponse(http.StatusOK, det)
	}

	// Is the job waiting to start
	if det, ok := cache.GetMetadata(jobId); ok {
		return MakeResponse(http.StatusOK, job.StorableDetails{
			UserID:   det.Owner,
			Token:    jobId,
			Status:   job.StatusNotStarted,
			Projects: det.Projects,
		})
	}

	// Is the job already completed
	if det, ok := cache.GetHistoricalDetails(jobId); ok {
		return MakeResponse(http.StatusOK, det)
	}

	// No job
	return svc.NotFound("no job found with id " + jobId)
}
