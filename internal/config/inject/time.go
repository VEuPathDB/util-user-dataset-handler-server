package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const timeInjectorTarget = "<<time>>"

// NewTimeInjector returns a new VariableInjector instance that will replace
// <<time>> variables in a command config.
func NewTimeInjector(
	details *job.Details,
	_ *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewTimeInjector")
	return &timeInjector{log, details}
}

type timeInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (t *timeInjector) Inject(target []string) ([]string, error) {
	t.log.Trace("inject.timeInjector.Inject")
	return simpleReplace(target, timeInjectorTarget,
		t.state.Started.Format("15:04:05")), nil
}
