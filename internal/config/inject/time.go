package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const timeInjectorTarget = "<<time>>"

func NewTimeInjector(details *job.Details, _ *job.Metadata) VariableInjector {
	return &timeInjector{details}
}

type timeInjector struct {
	state *job.Details
}

func (t *timeInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, timeInjectorTarget,
		t.state.Started.Format("15:04:05")), nil
}
