package workspace

import (
	"fmt"
	"io"
	"os"
	"path"

	"github.com/VEuPathDB/util-exporter-server/internal/util/fs"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/except"
)

const (
	workRoot = "/workspace"
	workPerm = 0666

	errDirCreateFail    = "failed to create workspace %s: %s"
	errTargetCreateFail = "failed to create local file for upload: %s"
	errUploadCopyFail   = "failed to copy uploaded file to workspace: %s"
	errOpenFail         = "failed to open file %s: %s"
	errStatFail         = "failed to stat file %s: %s"
	errDeleteFail       = "failed to delete file %s: %s"
)

type FilePredicate func(string) bool

type Workspace interface {
	// GetPath returns the path to the workspace directory.
	GetPath() string

	// FileFromUpload copies the contents of the given input stream into a new
	// file with the given name in the workspace directory.
	FileFromUpload(name string, in io.Reader) (*os.File, error)

	// Files returns a list of file names matching the given predicate
	// representing files in the workspace.
	Files(FilePredicate) ([]string, error)

	// Open the workspace file with the given name.
	Open(name string) (*os.File, error)

	// Delete the workspace file with the given name.
	Delete(name string) error

	// Stat returns the file info for the workspace file with the given name.
	Stat(name string) (os.FileInfo, error)
}

// Create returns a new Workspace instance rooted in the given directory path.
func Create(dir string, log *logrus.Entry) (Workspace, error) {
	log.Trace("Workspace#Create")

	root := path.Join(workRoot, dir)
	if err := os.MkdirAll(root, workPerm); err != nil {
		return nil, except.NewServerError(fmt.Sprintf(errDirCreateFail, root, err))
	}

	return &workspace{root, log}, nil
}

type workspace struct {
	dir string
	log *logrus.Entry
}

func (w *workspace) GetPath() string {
	return w.dir
}

func (w *workspace) Files(fn FilePredicate) ([]string, error) {
	tmp, err := fs.ListFiles(w.dir)

	if err != nil {
		return nil, except.NewServerError(err.Error())
	}

	out := make([]string, 0, len(tmp))

	for _, info := range tmp {
		if fn(info) {
			out = append(out, info)
		}
	}

	return out, nil
}

func (w *workspace) Open(name string) (*os.File, error) {
	fullPath := path.Join(w.dir, name)
	tmp, err := os.Open(fullPath)

	if err != nil {
		return nil, except.NewServerError(fmt.Sprintf(errOpenFail, fullPath, err))
	}

	return tmp, nil
}

func (w *workspace) Stat(name string) (os.FileInfo, error) {
	fullPath := path.Join(w.dir, name)
	tmp, err := os.Stat(fullPath)

	if err != nil {
		return nil, except.NewServerError(fmt.Sprintf(errStatFail, fullPath, err))
	}

	return tmp, nil
}

func (w *workspace) FileFromUpload(name string, in io.Reader) (*os.File, error) {
	w.log.Trace("Workspace#FileFromUpload")

	file, err := os.Create(path.Join(w.dir, name))

	if err != nil {
		return nil, except.NewServerError(fmt.Sprintf(errTargetCreateFail, err))
	}

	if _, err = io.Copy(file, in); err != nil {
		return nil, except.NewServerError(fmt.Sprintf(errUploadCopyFail, err))
	}

	return file, nil
}

func (w *workspace) Delete(name string) error {
	fullPath := path.Join(w.dir, name)

	if err := os.RemoveAll(fullPath); err != nil {
		return except.NewServerError(fmt.Sprintf(errDeleteFail, fullPath, err))
	}

	return nil
}
