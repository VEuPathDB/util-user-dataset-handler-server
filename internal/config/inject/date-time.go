package inject

import (
	"time"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dateTimeInjectTarget = "<<date-time>>"

// NewDateTimeInjector returns a new VariableInjector instance that will replace
// <<date-time>> variables in a command config.
func NewDateTimeInjector(det *job.Details, _ *job.Metadata) VariableInjector {
	return &dateTimeInjector{det}
}

type dateTimeInjector struct {
	state *job.Details
}

func (d *dateTimeInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dateTimeInjectTarget,
		d.state.Started.Format(time.RFC3339)), nil
}
