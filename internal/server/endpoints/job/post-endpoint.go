package job

import (
	"net/http"
	"os"
	"path"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/command"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/server/types"
)

const (
	tokenKey = "token"
)

func NewUploadEndpoint(o *config.Options, meta, upload *cache.Cache) types.Endpoint {
	return &endpoint{
		opt:    o,
		meta:   meta,
		upload: upload,
	}
}

type endpoint struct {
	log    *logrus.Entry
	opt    *config.Options
	meta   *cache.Cache
	upload *cache.Cache
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
// If we've made it this far we know that the
func (e *endpoint) Handle(req midl.Request) midl.Response {
	token := mux.Vars(req.RawRequest())[tokenKey]
	meta  := e.getMeta(token)
	dets  := e.CreateDetails(&meta)
	log   := middle.GetCtxLogger(req)

	if res := e.HandleUpload(req, dets); res != nil {
		return res
	}

	stream, err := command.NewCommandRunner(token, e.opt, e.upload, e.meta, log).
		Run()
	if err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		return svc.ServerError(err.Error())
	}

	return midl.MakeResponse(http.StatusOK, stream).
		Callback(e.cleanup(token))
}

func (e *endpoint) HandleUpload(
	request midl.Request,
	details *job.Details,
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
		return e.FailJob(err, details)
	}
	defer out.Close()

	if err := CopyFile(out, upload, log); err != nil {
		return e.FailJob(err, details)
	}

	details.InTarName = head.Filename
	e.StoreDetails(details)

	return nil
}

func (e *endpoint) getMeta(token string) Metadata {
	tmp, _ := e.meta.Get(token)
	return tmp.(Metadata)
}

func (e *endpoint) cleanup(token string) func() {
	return func() {
		tmp, _ := e.upload.Get(token)
		dets := tmp.(job.Details)

		_ = os.RemoveAll(dets.WorkingDir)
		e.upload.Set(token, dets.StorableDetails, cache.DefaultExpiration)
	}
}