package inject

import (
	"strconv"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const timestampInjectorTarget = "<<timestamp>>"

// NewTimestampInjector returns a new VariableInjector instance that will
// replace <<timestamp>> variables in a command config.
func NewTimestampInjector(details *job.Details, _ *job.Metadata) VariableInjector {
	return &timestampInjector{details}
}

type timestampInjector struct {
	state *job.Details
}

func (t *timestampInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, timestampInjectorTarget,
		strconv.FormatInt(t.state.Started.Unix(), 10)), nil
}
