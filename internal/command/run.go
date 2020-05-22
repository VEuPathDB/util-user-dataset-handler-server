package command

import (
	// Std lib
	"errors"
	"fmt"
	"io"
	"os"
	"path"
	"strings"
	"time"

	// External
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	errComFail = `Command "%s" failed: %s`
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
	var stat job.StatusFile
	var err  error
	var pack []string

	r.getDetails()
	r.updateStatus(job.StatusUnpacking)

	if err = r.unpack(&r.details); err != nil {
		return r.fail(err)
	}

	r.updateStatus(job.StatusProcessing)
	for i := range r.options.Commands {
		if err = r.handleCommand(&r.options.Commands[i]); err != nil {
			return r.fail(err)
		}
		if stat, err = r.getStatusFile(); err != nil {
			return r.fail(err)
		} else {
			if strings.ToLower(stat.Status) != "ok" {
				return r.fail(fmt.Errorf(errComFail, r.options.Commands[i].Command,
					stat.Message))
			}
			r.details.OutputFiles = append(r.details.OutputFiles, stat.Files)
			pack = append(pack, stat.Pack...)
		}
	}

	pack, err = r.generateMetaFiles(pack)
	if err != nil {
		return r.fail(err)
	}

	r.updateStatus(job.StatusPacking)
	err = r.packArchive(pack)

	file, err :=  os.OpenFile(path.Join(r.details.WorkingDir, "dataset.tgz"),
		os.O_RDONLY, dsFilePerm)
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