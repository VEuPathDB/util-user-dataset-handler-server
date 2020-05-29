package command

import (
	// Std lib
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"

	// External
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
	"github.com/VEuPathDB/util-exporter-server/internal/service/workspace"
)

type RunResult struct {
	Error  error
	Stream io.ReadCloser
	Name   string
}

func NewCommandRunner(
	token string,
	options *config.Options,
	wkspc workspace.Workspace,
	log *logrus.Entry,
) Runner {
	return &runner{
		log:     log,
		token:   token,
		options: options,
		wkspc:   wkspc,
	}
}

type Runner interface {
	Run() RunResult
}

type runner struct {
	log     *logrus.Entry
	token   string
	options *config.Options

	lastStatus  time.Time
	lastCommand *config.Command
	details     job.Details
	meta        job.Metadata
	wkspc       workspace.Workspace
}

func (r *runner) Run() RunResult {
	var err error

	r.getDetails()
	r.getMeta()
	r.updateStatus(job.StatusUnpacking)

	if err = r.unpack(&r.details); err != nil {
		return r.fail(err)
	}

	r.updateStatus(job.StatusProcessing)

	if err = r.handleCommand(&r.options.Command); err != nil {
		return r.fail(err)
	}

	fileName, err := r.findTar()
	if err != nil {
		return r.fail(err)
	}

	file, err := r.wkspc.Open(fileName)
	if err != nil {
		return r.fail(
			errors.New("Failed to open packaged tar for reading: " + err.Error()))
	}

	r.updateStatus(job.StatusCompleted)

	return RunResult{
		Stream: file,
		Name:   fileName,
	}
}

func (r *runner) getDetails() {
	r.details, _ = cache.GetDetails(r.token)
}

func (r *runner) storeDetails() {
	cache.PutDetails(r.token, r.details)
}

func (r *runner) getMeta() {
	r.meta, _ = cache.GetMetadata(r.token)
}

func (r *runner) updateStatus(status job.Status) {
	r.details.Status = status
	r.storeDetails()
}

func (r *runner) findTar() (string, error) {
	prefix := fmt.Sprintf("dataset_u%d", r.meta.Owner)

	matches, err := r.wkspc.Files(func(f os.FileInfo) bool {
		return strings.HasPrefix(f.Name(), prefix)
	})

	if err != nil {
		return "", err
	}

	switch len(matches) {
	case 0:
		return "", errors.New("no dataset archive found")
	case 1:
		return matches[0].Name(), nil
	default:
		return "", errors.New("invalid state, more than one dataset archive present in workspace")
	}
}
