package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

const dateInjectorTarget = "<<date>>"

func NewDateInjector(det *process.Details) VariableInjector {
	return &dateInjector{det}
}

type dateInjector struct {
	state *process.Details
}

func (d *dateInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dateInjectorTarget,
		d.state.Started.Format("2006-01-02")), nil
}
