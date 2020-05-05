package command

import (
	// Std lib
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/server/endpoints/metadata"
	"github.com/sirupsen/logrus"
	"io"
	"strings"
	"time"

	// External
	"github.com/patrickmn/go-cache"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

const (
	errComFail = `Command "%s" failed: %s`
)

func NewCommandRunner(
	token string,
	options *config.Options,
	upload *cache.Cache,
	meta *cache.Cache,
) Runner {
	return &runner{
		token:   token,
		options: options,
		upload:  upload,
		meta:    meta,
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
	details     process.Details
}

func (r *runner) Run() (io.ReadCloser, error) {
	var stat process.StatusFile
	var err  error
	var pack []string

	r.getDetails()
	r.updateStatus(process.StatusUnpacking)

	if err = r.unpack(&r.details); err != nil {
		return r.fail(err)
	}

	r.updateStatus(process.StatusProcessing)
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

	r.updateStatus(process.StatusPacking)
	err = r.packArchive(pack)

}

func (r *runner) getDetails() {
	tmp, _ := r.upload.Get(r.token)
	r.details = tmp.(process.Details)
}

func (r *runner) storeDetails() {
	r.upload.Set(r.token, r.details, cache.DefaultExpiration)
}

func (r *runner) getMeta() metadata.Metadata {
	tmp, _ := r.upload.Get(r.token)
	return tmp.(metadata.Metadata)
}

func (r *runner) updateStatus(status process.Status) {
	r.details.Status = status
	r.storeDetails()
}