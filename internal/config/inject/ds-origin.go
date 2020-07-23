package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const dsOriginInjectorTarget = "<<ds-origin>>"

// NewDsNameInjector returns a new VariableInjector instance that will replace
// <<ds-name>> variables in a command config.
func NewDsOriginInjector(
	_ *job.Details,
	meta *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewDsOriginInjector")
	return &dsOriginInjector{log, meta}
}

type dsOriginInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (d *dsOriginInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dsOriginInjector.Inject")
	return simpleReplace(target, dsOriginInjectorTarget,
		d.state.Origin), nil
}
