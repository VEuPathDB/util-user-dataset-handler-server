package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const cwdInjectorTarget = "<<cwd>>"

func NewCwdInjector(det *job.Details) VariableInjector {
	return &cwdInjector{det}
}

type cwdInjector struct {
	state *job.Details
}

func (d *cwdInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, cwdInjectorTarget,
		d.state.WorkingDir), nil
}
