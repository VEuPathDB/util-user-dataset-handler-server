package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/except"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"os"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/command"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
	"github.com/VEuPathDB/util-exporter-server/internal/service/workspace"
)

var (
	promRequestPayloadSize = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "upload",
		Name:      "file_size",
		Help:      "File size for user uploads in MiB.",
		Buckets: []float64{
			0.5,  // 0.5 MiB
			1,    // 1   MiB
			10,   // 10  MiB
			50,   // 50  MiB
			100,  // 100 MiB
			250,  // 250 MiB
			500,  // 500 MiB
			1024, // 1   GiB
		},
	}, []string{"ext"})
)

// NewUploadEndpoint instantiates a new endpoint wrapper for the user dataset
// upload handler.
func NewUploadEndpoint(opts *config.Options) types.Endpoint {
	return &endpoint{opt: opts}
}

type endpoint struct {
	log    *logrus.Entry
	opt    *config.Options
}

func (e *endpoint) Register(r *mux.Router) {
	r.Path(urlPath).
		Methods(http.MethodPost).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(
			middle.NewBinaryAdaptor().AddHandlers(
				middle.NewTokenValidator(tokenKey, e)))))
}

// Handle the request.
//
// If we've made it this far we know that the token in the URL is valid and
// points to an existing metadata entry in the store.
func (e *endpoint) Handle(req midl.Request) midl.Response {
	token := mux.Vars(req.RawRequest())[tokenKey]
	meta  := e.getMeta(token)
	dets  := e.CreateDetails(&meta)
	log   := logger.ByRequest(req)
	e.log = log

	wkspc, err := workspace.Create(token, log)
	if err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		return svc.ServerError(err.Error())
	}

	if res := e.HandleUpload(req, dets, wkspc); res != nil {
		return res
	}

	result := command.NewCommandRunner(token, e.opt, wkspc, log).Run()
	if result.Error != nil {
		switch result.Error.(type) {
		case *command.UserError:
			log.WithField("status", http.StatusUnprocessableEntity).Error(result.Error)
			return svc.InvalidRequest(result.Error.Error()).Callback(e.cleanup(token))
		default:
			log.WithField("status", http.StatusInternalServerError).Error(result.Error)
			return svc.ServerError(result.Error.Error()).Callback(e.cleanup(token))
		}
	}

	return midl.MakeResponse(http.StatusOK, result).Callback(e.cleanup(token))
}

func (e *endpoint) HandleUpload(
	request midl.Request,
	details *job.Details,
	wkspc workspace.Workspace,
) midl.Response {
	log := e.log

	upload, head, res := GetFileHandle(request.RawRequest(), log)
	if res != nil {
		return e.FailJob(res, details)
	}
	defer upload.Close()

	suff, errRes := ValidateFileSuffix(head.Filename, log)

	if errRes != nil {
		return e.FailJob(errRes, details)
	}

	details.WorkingDir = wkspc.GetPath()
	e.StoreDetails(details)

	file, err := wkspc.FileFromStream(head.Filename, upload)
	if err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		return svc.ServerError(err.Error())
	}
	defer file.Close()

	if info, err := file.Stat(); err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		return svc.ServerError(except.NewServerError(err.Error()).Error())
	} else {
		promRequestPayloadSize.WithLabelValues(suff).
			Observe(float64(info.Size()) / float64(util.SizeMebibyte))
	}

	details.InTarName = head.Filename
	e.StoreDetails(details)

	return nil
}

// retrieve metadata from the metadata store
func (e *endpoint) getMeta(token string) job.Metadata {
	tmp, _ := cache.GetMetadata(token)
	return tmp
}

// remove the working directory and convert the stored metadata to the long
// store form.
func (e *endpoint) cleanup(token string) func() {
	return func() {
		e.log.Debug("cleaning up workspace")
		details, _ := cache.GetDetails(token)

		_ = os.RemoveAll(details.WorkingDir)
		cache.PutHistoricalDetails(token, details.StorableDetails)
		cache.DeleteMetadata(token)
		cache.DeleteDetails(token)
	}
}