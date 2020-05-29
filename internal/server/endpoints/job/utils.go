package job

import (
	"mime/multipart"
	"net/http"
	"strings"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

var allowedSuffixes = []string{
	".tar",
	".zip",
	".tgz",
	".tar.gz",
}

const (
	errBadSuffix = "Invalid file format, must be one of: "
)

func ValidateFileSuffix(file string, log *logrus.Entry) (string, midl.Response) {
	for i := range allowedSuffixes {
		if strings.HasSuffix(file, allowedSuffixes[i]) {
			return allowedSuffixes[i], nil
		}
	}

	msg := errBadSuffix + strings.Join(allowedSuffixes, ", ")
	log.WithField("status", http.StatusBadRequest).Info(msg)

	return "", svc.BadRequest(msg)
}

const (
	errNoBoundary = "Malformed request, no form boundary."
	errNoFile     = "No file was provided at the expected key 'file'."
	errNotMulti   = "Malformed request, server expects valid multipart/form-data"
	errUnknown    = "Invalid request: "
)

func GetFileHandle(req *http.Request, log *logrus.Entry) (
	file multipart.File,
	head *multipart.FileHeader,
	out midl.Response,
) {
	file, head, err := req.FormFile("file")

	if err != nil {
		switch err {
		case http.ErrMissingBoundary:
			log.WithField("status", http.StatusBadRequest).Info(errNoBoundary)

			out = svc.BadRequest(errNoBoundary)
		case http.ErrMissingFile:
			log.WithField("status", http.StatusBadRequest).Info(errNoFile)

			out = svc.BadRequest(errNoFile)
		case http.ErrNotMultipart:
			log.WithField("status", http.StatusBadRequest).Info(errNotMulti)

			out = svc.BadRequest(errNotMulti)
		default:
			log.WithField("status", http.StatusBadRequest).
				Info(errUnknown + err.Error())

			out = svc.BadRequest(errUnknown + err.Error())
		}
	}

	return
}

func (e *endpoint) FailJob(out midl.Response, details *job.Details) midl.Response {
	details.Status = job.StatusFailed
	e.storeDetails(details)
	out.Callback(e.cleanup(details.Token))

	return out
}
