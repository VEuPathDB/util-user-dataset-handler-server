package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsOriginInjectorTarget = "<<ds-origin>>"

// NewDsNameInjector returns a new VariableInjector instance that will replace
// <<ds-name>> variables in a command config.
func NewDsOriginInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsOriginInjector{meta}
}

type dsOriginInjector struct {
	state *job.Metadata
}

func (d *dsOriginInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsOriginInjectorTarget,
		d.state.Origin), nil
}
