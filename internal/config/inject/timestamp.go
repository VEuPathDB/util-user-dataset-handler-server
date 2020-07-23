package inject

import (
	"github.com/sirupsen/logrus"
	"strconv"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const timestampInjectorTarget = "<<timestamp>>"

// NewTimestampInjector returns a new VariableInjector instance that will
// replace <<timestamp>> variables in a command config.
func NewTimestampInjector(
	details *job.Details,
	_ *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewTimestampInjector")
	return &timestampInjector{log, details}
}

type timestampInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (t *timestampInjector) Inject(target []string) ([]string, error) {
	t.log.Trace("inject.timestampInjector.Inject")
	return simpleReplace(target, timestampInjectorTarget,
		strconv.FormatInt(t.state.Started.Unix(), 10)), nil
}
