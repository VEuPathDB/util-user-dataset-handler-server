package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const dateInjectorTarget = "<<date>>"

// NewDateInjector returns a new VariableInjector instance that will replace
// <<date>> variables in a command config.
func NewDateInjector(det *job.Details, _ *job.Metadata, log *logrus.Entry) VariableInjector {
	log.Trace("inject.NewDateInjector")
	return &dateInjector{log, det}
}

type dateInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (d *dateInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dateInjector.Inject")
	return simpleReplace(target, dateInjectorTarget,
		d.state.Started.Format("2006-01-02")), nil
}
