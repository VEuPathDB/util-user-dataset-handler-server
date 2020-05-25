package command

import (
	// Std lib
	"errors"
	"io"
	"os"
	"path"
	"time"

	// External
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func NewCommandRunner(
	token   string,
	options *config.Options,
	upload  *cache.Cache,
	meta    *cache.Cache,
	ctxLog  *logrus.Entry,
) Runner {
	return &runner{
		token:   token,
		options: options,
		upload:  upload,
		meta:    meta,
		log:     ctxLog,
	}
}

type Runner interface {
	Run() (io.ReadCloser, error)
}

type runner struct {
	log     *logrus.Entry
	token   string
	options *config.Options
	upload  *cache.Cache
	meta    *cache.Cache

	lastStatus  time.Time
	lastCommand *config.Command
	details     job.Details
}

func (r *runner) Run() (io.ReadCloser, error) {
	var err  error

	r.getDetails()
	r.updateStatus(job.StatusUnpacking)

	if err = r.unpack(&r.details); err != nil {
		return r.fail(err)
	}

	r.updateStatus(job.StatusProcessing)

	if err = r.handleCommand(&r.options.Command); err != nil {
		return r.fail(err)
	}

	r.updateStatus(job.StatusPacking)

	if err = r.packArchive(); err != nil {
		return r.fail(err)
	}

	file, err :=  os.Open(path.Join(r.details.WorkingDir, "dataset.tgz"))
	if err != nil {
		return r.fail(
			errors.New("Failed to open packaged tar for reading: " + err.Error()))
	}

	return file, nil
}

func (r *runner) getDetails() {
	tmp, _ := r.upload.Get(r.token)
	r.details = tmp.(job.Details)
}

func (r *runner) storeDetails() {
	r.upload.Set(r.token, r.details, cache.DefaultExpiration)
}

func (r *runner) getMeta() job.Metadata {
	tmp, _ := r.upload.Get(r.token)
	return tmp.(job.Metadata)
}

func (r *runner) updateStatus(status job.Status) {
	r.details.Status = status
	r.storeDetails()
}