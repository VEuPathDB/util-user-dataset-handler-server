package inject

import (
	"github.com/sirupsen/logrus"
	"time"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dateTimeInjectTarget = "<<date-time>>"

// NewDateTimeInjector returns a new VariableInjector instance that will replace
// <<date-time>> variables in a command config.
func NewDateTimeInjector(det *job.Details, _ *job.Metadata, log *logrus.Entry) VariableInjector {
	log.Trace("inject.NewDateTimeInjector")
	return &dateTimeInjector{log, det}
}

type dateTimeInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (d *dateTimeInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dateTimeInjector.Inject")
	return simpleReplace(target, dateTimeInjectTarget,
		d.state.Started.Format(time.RFC3339)), nil
}
