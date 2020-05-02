package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"time"
)

const dateTimeInjectTarget = "<<date-time>>"

func NewDateTimeInjector(det *process.Details) VariableInjector {
	return &dateTimeInjector{det}
}

type dateTimeInjector struct {
	state *process.Details
}

func (d *dateTimeInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dateTimeInjectTarget,
		d.state.Started.Format(time.RFC3339)), nil
}

