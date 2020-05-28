package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsNameInjectorTarget = "<<ds-name>>"

func NewDsNameInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsNameInjector{meta}
}

type dsNameInjector struct {
	state *job.Metadata
}

func (d *dsNameInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsNameInjectorTarget,
		d.state.Name), nil
}
