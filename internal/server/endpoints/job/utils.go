package job

import (
	"io"
	"net/http"
	"strings"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/sirupsen/logrus"
	"github.com/vulpine-io/bites/v1/pkg/bites"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

const (
	errBadSuffix = "Invalid file format, must be one of: "
)

func (u *uploadEndpoint) ValidateFileSuffix(file string, log *logrus.Entry) (string, midl.Response) {
	for _, s := range u.file.FileTypes() {
		if strings.HasSuffix(file, s) {
			return s, nil
		}
	}

	msg := errBadSuffix + strings.Join(u.file.FileTypes(), ", ")
	log.WithField("status", http.StatusBadRequest).Info(msg)

	return "", svc.BadRequest(msg)
}

const (
	errNoBoundary = "Malformed request, no form boundary."
	errNoFile     = "No file was provided at the expected key 'file'."
	errNotMulti   = "Malformed request, server expects valid multipart/form-data"
	errUnknown    = "Error: "
)

func GetFileHandle(req *http.Request, log *logrus.Entry) (
	fileName string,
	stream io.ReadCloser,
	out midl.Response,
) {
	log.Trace("job.GetFileHandle")

	req.Body = http.MaxBytesReader(nil, req.Body, int64(bites.SizeGibibyte))

	reader, err := req.MultipartReader()
	if err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		out = svc.ServerError(err.Error())
		return
	}

	part, err := reader.NextPart()
	if err != nil {
		if err == io.EOF {
			log.WithField("status", http.StatusBadRequest).Info("empty form data")
			out = svc.BadRequest("empty form data body")
			return
		} else {
			log.WithField("status", http.StatusInternalServerError).Error(err)
			out = svc.ServerError(err.Error())
			return
		}
	}

	if part.FormName() != "file" {
		msg := "invalid form body.  expected single part with name \"file\""
		log.WithField("status", http.StatusBadRequest).Info(msg)
		out = svc.ServerError(msg)
		return
	}

	fileName = part.FileName()
	stream = part

	return
}

func (u *uploadEndpoint) FailJob(out midl.Response, details *job.Details) midl.Response {
	details.Status = job.StatusFailed
	u.storeDetails(details)
	out.Callback(u.cleanup(details.Token))

	return out
}
