package job

import (
	"net/http"
	"os"
	"path"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/command"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
)

// NewUploadEndpoint instantiates a new endpoint wrapper for the user dataset
// upload handler.
func NewUploadEndpoint(o *config.Options, meta *cache.Meta, upload *cache.Upload) types.Endpoint {
	return &endpoint{
		opt:    o,
		meta:   meta,
		upload: upload,
	}
}

type endpoint struct {
	log    *logrus.Entry
	opt    *config.Options
	meta   *cache.Meta
	upload *cache.Upload
}

func (e *endpoint) Register(r *mux.Router) {
	r.Path(urlPath).Handler(middle.NewBinaryAdaptor().
		AddHandlers(
			middle.RequestIdProvider(),
			middle.LogProvider(),
			middle.NewTimer(middle.NewTokenValidator(tokenKey, e.meta, e))))
}

// Handle the request.
//
// If we've made it this far we know that the token in the URL is valid and
// points to an existing metadata entry in the store.
func (e *endpoint) Handle(req midl.Request) midl.Response {
	token := mux.Vars(req.RawRequest())[tokenKey]
	meta  := e.getMeta(token)
	dets  := e.CreateDetails(&meta)
	log   := middle.GetCtxLogger(req)
	e.log = log

	if res := e.HandleUpload(req, dets, meta); res != nil {
		return res
	}

	result := command.NewCommandRunner(token, e.opt, e.upload, e.meta, log).Run()
	if result.Error != nil {
		log.WithField("status", http.StatusInternalServerError).Error(result.Error)
		return svc.ServerError(result.Error.Error()).Callback(e.cleanup(token))
	}

	return midl.MakeResponse(http.StatusOK, result).Callback(e.cleanup(token))
}

func (e *endpoint) HandleUpload(
	request midl.Request,
	details *job.Details,
	meta job.Metadata,
) midl.Response {
	log := middle.GetCtxLogger(request)

	upload, head, res := GetFileHandle(request.RawRequest(), log)
	if res != nil {
		return e.FailJob(res, details)
	}
	defer upload.Close()

	if err := ValidateFileSuffix(head.Filename, log); err != nil {
		return e.FailJob(err, details)
	}

	dir := path.Join(e.opt.Workspace, details.Token)
	if err := MakeWorkDir(dir, log); err != nil {
		return e.FailJob(err, details)
	}

	details.WorkingDir = dir
	e.StoreDetails(details)

	filename := path.Join(dir, head.Filename)
	out, err := MakeFileTarget(filename, log)
	if err != nil {
		return e.FailJob(err, details).Callback(e.cleanup(meta.Token))
	}
	defer out.Close()

	if err := CopyFile(out, upload, log); err != nil {
		return e.FailJob(err, details).Callback(e.cleanup(meta.Token))
	}

	details.InTarName = head.Filename
	e.StoreDetails(details)

	return nil
}

// retrieve metadata from the metadata store
func (e *endpoint) getMeta(token string) job.Metadata {
	tmp, _ := e.meta.Get(token)
	return tmp
}

// remove the working directory and convert the stored metadata to the long
// store form.
func (e *endpoint) cleanup(token string) func() {
	return func() {
		e.log.Debug("cleaning up workspace")
		dets, _ := e.upload.GetDetails(token)

		_ = os.RemoveAll(dets.WorkingDir)
		e.upload.SetStorable(token, dets.StorableDetails)
	}
}