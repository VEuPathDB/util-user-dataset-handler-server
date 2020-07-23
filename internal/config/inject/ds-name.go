package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const dsNameInjectorTarget = "<<ds-name>>"

// NewDsNameInjector returns a new VariableInjector instance that will replace
// <<ds-name>> variables in a command config.
func NewDsNameInjector(
	_ *job.Details,
	meta *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewDsNameInjector")
	return &dsNameInjector{log, meta}
}

type dsNameInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (d *dsNameInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dsNameInjector.Inject")
	return forceQuotesReplace(target, dsNameInjectorTarget,
		d.state.Name), nil
}
