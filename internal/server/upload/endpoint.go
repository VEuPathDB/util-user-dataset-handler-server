package upload

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/command"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"github.com/VEuPathDB/util-exporter-server/internal/server/metadata"
	"github.com/gorilla/mux"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
	"path"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

const (
	tokenKey = "token"
	urlPath  = "/process/dataset/{" + tokenKey + "}"
)

const (
	errNoMeta = "Invalid state, missing metadata"
)

func NewUploadEndpoint(o *config.Options, meta, upload *cache.Cache) svc.Endpoint {
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
			middle.NewContentLengthFilter(500 * util.SizeMebibyte),
			middle.NewLogProvider(NewTokenFilter(e.meta)),
			middle.NewLogProvider(middle.NewTimer(
				func(logger *logrus.Entry) midl.Middleware {
					e.log = logger
					return e
				}))))
}

// Handle the request.
//
// If we've made it this far we know that the
func (e *endpoint) Handle(req midl.Request) midl.Response {
	token := mux.Vars(req.RawRequest())[tokenKey]
	meta  := e.getMeta(token)
	dets  := e.CreateDetails(&meta)

	e.HandleUpload(req, dets)

	stream, err := command.NewCommandRunner(token, e.opt, e.upload, e.meta).Run()
	if err != nil {
		return svc.ServerError(err.Error())
	}

	return midl.MakeResponse(http.StatusOK, stream).
		Callback(e.cleanup(token))
}

func (e *endpoint) HandleUpload(
	request midl.Request,
	details *process.Details,
) midl.Response {

	upload, head, res := GetFileHandle(request.RawRequest())
	if res != nil {
		return e.FailJob(res, details)
	}
	defer upload.Close()

	if err := ValidateFileSuffix(head.Filename); err != nil {
		return e.FailJob(err, details)
	}

	dir := path.Join(e.opt.Workspace, details.Token)
	if err := MakeWorkDir(dir); err != nil {
		return e.FailJob(err, details)
	}

	details.WorkingDir = dir
	e.StoreDetails(details)

	filename := path.Join(dir, head.Filename)
	out, err := MakeFileTarget(filename)
	if err != nil {
		return e.FailJob(err, details)
	}
	defer out.Close()

	if err := CopyFile(out, upload); err != nil {
		return e.FailJob(err, details)
	}

	details.InTarName = head.Filename
	e.StoreDetails(details)

	return nil
}

func (e *endpoint) getMeta(token string) metadata.Metadata {
	tmp, _ := e.meta.Get(token)
	return tmp.(metadata.Metadata)
}

func (e *endpoint) cleanup(token string) func() {
	return func() {
		tmp, _ := e.upload.Get(token)
		dets := tmp.(process.Details)

		_ = os.RemoveAll(dets.WorkingDir)
		e.upload.Set(token, dets.StorableDetails, cache.DefaultExpiration)
	}
}