package job

import (
	"github.com/sirupsen/logrus"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"

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

func ValidateFileSuffix(file string, log *logrus.Entry) midl.Response {
	for i := range allowedSuffixes {
		if strings.HasSuffix(file, allowedSuffixes[i]) {
			return nil
		}
	}

	msg := errBadSuffix + strings.Join(allowedSuffixes, ", ")
	log.WithField("status", http.StatusBadRequest).Info(msg)
	return svc.BadRequest(msg)
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
	out  midl.Response,
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

const (
	errDirFail  = "Could not create workspace directory "
	errFileFail = "Could not open target file for writing: "
	errCopyFail = "Could not copy upload to workspace directory: "
)

// MakeWorkDir attempts to create a working directory at the path given.
func MakeWorkDir(dir string, log *logrus.Entry) midl.Response {
	if err := os.MkdirAll(dir, 0666); err != nil {
		log.WithField("status", http.StatusInternalServerError).
			Error(errDirFail + dir)
		return svc.ServerError(errDirFail + dir)
	}
	return nil
}

// MakeFileTarget attempts to create a file in the working directory in which
// the user input will be saved.
func MakeFileTarget(file string, log *logrus.Entry) (*os.File, midl.Response) {
	out, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		log.WithField("status", http.StatusInternalServerError).Error(err)
		return nil, svc.ServerError(errFileFail + err.Error())
	}
	return out, nil
}

// CopyFile attempts to copy the input stream from the user request to the
// given output stream.
func CopyFile(out io.Writer, in io.Reader, log *logrus.Entry) midl.Response {
	if _, err := io.Copy(out, in); err != nil {
		log.WithField("status", http.StatusInternalServerError).
			Error(errCopyFail + err.Error())
		return svc.ServerError(errCopyFail + err.Error())
	}
	return nil
}


func (e *endpoint) FailJob(out midl.Response, details *job.Details) midl.Response {
	details.Status = job.StatusFailed
	e.StoreDetails(details)
	out.Callback(e.cleanup(details.Token))
	return out
}