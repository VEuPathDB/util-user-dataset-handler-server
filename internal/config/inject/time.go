package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

const timeInjectorTarget = "<<time>>"

func NewTimeInjector(details *process.Details) VariableInjector {
	return &timeInjector{details}
}

type timeInjector struct {
	state *process.Details
}

func (t *timeInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, timeInjectorTarget,
		t.state.Started.Format("15:04:05")), nil
}
