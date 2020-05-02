package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

const cwdInjectorTarget = "<<cwd>>"

func NewCwdInjector(det *process.Details) VariableInjector {
	return &cwdInjector{det}
}

type cwdInjector struct {
	state *process.Details
}

func (d *cwdInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, cwdInjectorTarget,
		d.state.WorkingDir), nil
}
