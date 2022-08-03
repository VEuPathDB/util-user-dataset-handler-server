package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const dsDescriptionInjectorTarget = "<<ds-description>>"

// NewDsDescriptionInjector returns a new VariableInjector instance that will
// replace <<ds-description>> variables in a command config.
func NewDsDescriptionInjector(_ *job.Details, meta *job.Metadata, log *logrus.Entry) VariableInjector {
	log.Trace("inject.NewDsDescriptionInjector")
	return &dsDescriptionInjector{log, meta}
}

type dsDescriptionInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (d *dsDescriptionInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dsDescriptionInjector.Inject")
	return simpleReplace(target, dsDescriptionInjectorTarget,
		d.state.Description), nil
}
