package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"strconv"
)

const timestampInjectorTarget = "<<timestamp>>"

func NewTimestampInjector(details *process.Details) VariableInjector {
	return &timestampInjector{details}
}

type timestampInjector struct {
	state *process.Details
}

func (t *timestampInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, timestampInjectorTarget,
		strconv.FormatInt(t.state.Started.Unix(), 10)), nil
}

