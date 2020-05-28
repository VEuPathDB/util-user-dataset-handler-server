package command

import (
	// Std lib
	"errors"
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/cache"
	"io"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"time"

	// External
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

type RunResult struct {
	Error  error
	Stream io.ReadCloser
	Name   string
}

func NewCommandRunner(
	token string,
	options *config.Options,
	upload *cache.Upload,
	meta *cache.Meta,
	ctxLog *logrus.Entry,
) Runner {
	return &runner{
		token:       token,
		options:     options,
		uploadCache: upload,
		metaCache:   meta,
		log:         ctxLog,
	}
}

type Runner interface {
	Run() RunResult
}

type runner struct {
	log         *logrus.Entry
	token       string
	options     *config.Options
	uploadCache *cache.Upload
	metaCache   *cache.Meta

	lastStatus  time.Time
	lastCommand *config.Command
	details     job.Details
	meta        job.Metadata
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

	file, err := os.Open(path.Join(r.details.WorkingDir, fileName))
	if err != nil {
		return r.fail(
			errors.New("Failed to open packaged tar for reading: " + err.Error()))
	}

	r.updateStatus(job.StatusCompleted)

	return RunResult{
		Stream: file,
		Name: fileName,
	}
}

func (r *runner) getDetails() {
	r.details, _ = r.uploadCache.GetDetails(r.token)
}

func (r *runner) storeDetails() {
	r.uploadCache.SetDetails(r.token, r.details)
}

func (r *runner) getMeta() {
	r.meta, _ = r.metaCache.Get(r.token)
}

func (r *runner) updateStatus(status job.Status) {
	r.details.Status = status
	r.storeDetails()
}

func (r *runner) findTar() (string, error) {
	files, err := ioutil.ReadDir(r.details.WorkingDir)

	if err != nil {
		return "", errors.New(errDirReadFail + err.Error())
	}

	prefix := fmt.Sprintf("dataset_u%d", r.meta.Owner)

	for _, f := range files {
		if strings.HasPrefix(f.Name(), prefix) {
			return f.Name(), nil
		}
	}

	return "", errors.New("no dataset archive found")
}
