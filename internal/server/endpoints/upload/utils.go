package upload

import (
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strings"
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

func ValidateFileSuffix(file string) midl.Response {
	for i := range allowedSuffixes {
		if strings.HasSuffix(file, allowedSuffixes[i]) {
			return nil
		}
	}
	return svc.BadRequest(errBadSuffix + strings.Join(allowedSuffixes, ", "))
}

const (
	errNoBoundary = "Malformed request, no form boundary."
	errNoFile     = "No file was provided at the expected key 'file'."
	errNotMulti   = "Malformed request, server expects valid multipart/form-data"
	errUnknown    = "Invalid request: "
)

func GetFileHandle(req *http.Request) (
	file multipart.File,
	head *multipart.FileHeader,
	out midl.Response,
) {
	file, head, err := req.FormFile("file")

	if err != nil {
		switch err {
		case http.ErrMissingBoundary:
			out = svc.BadRequest(errNoBoundary)
		case http.ErrMissingFile:
			out = svc.BadRequest(errNoFile)
		case http.ErrNotMultipart:
			out = svc.BadRequest(errNotMulti)
		default:
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

func MakeWorkDir(dir string) midl.Response {
	if err := os.MkdirAll(dir, 0666); err != nil {
		return svc.ServerError(errDirFail + dir)
	}
	return nil
}

func MakeFileTarget(file string) (*os.File, midl.Response) {
	out, err := os.OpenFile(file, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, svc.ServerError(errFileFail + err.Error())
	}
	return out, nil
}

func CopyFile(out io.Writer, in io.Reader) midl.Response {
	if _, err := io.Copy(out, in); err != nil {
		return svc.ServerError(errCopyFail + err.Error())
	}
	return nil
}


func (e *endpoint) FailJob(out midl.Response, details *process.Details) midl.Response {
	details.Status = process.StatusFailed
	e.StoreDetails(details)
	out.Callback(e.cleanup(details.Token))
	return out
}