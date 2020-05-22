package command

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"io"
)

func (r *runner) fail(err error) (io.ReadCloser, error) {
	r.details.Status = job.StatusFailed
	r.storeDetails()
	r.log.Error(err)
	return nil, err
}
