package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const cwdInjectorTarget = "<<cwd>>"

// NewCwdInjector returns a new VariableInjector instance that will replace
// <<cwd>> variables in a command config.
func NewCwdInjector(det *job.Details, _ *job.Metadata) VariableInjector {
	return &cwdInjector{det}
}

type cwdInjector struct {
	state *job.Details
}

func (d *cwdInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, cwdInjectorTarget,
		d.state.WorkingDir), nil
}
