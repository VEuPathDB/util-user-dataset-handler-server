package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsNameInjectorTarget = "<<ds-name>>"

// NewDsNameInjector returns a new VariableInjector instance that will replace
// <<ds-name>> variables in a command config.
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
