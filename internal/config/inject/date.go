package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dateInjectorTarget = "<<date>>"

func NewDateInjector(det *job.Details, _ *job.Metadata) VariableInjector {
	return &dateInjector{det}
}

type dateInjector struct {
	state *job.Details
}

func (d *dateInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dateInjectorTarget,
		d.state.Started.Format("2006-01-02")), nil
}
