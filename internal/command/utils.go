package command

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func (r *runner) fail(err error) RunResult {
	r.log.Trace("command.runner.fail")

	r.details.Status = job.StatusFailed
	r.storeDetails()
	r.log.Error(err)

	return RunResult{Error: err}
}
