package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsSummaryInjectorTarget = "<<ds-summary>>"

// NewDsSummaryInjector returns a new VariableInjector instance that will
// replace <<ds-summary>> variables in a command config.
func NewDsSummaryInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsSummaryInjector{meta}
}

type dsSummaryInjector struct {
	state *job.Metadata
}

func (d *dsSummaryInjector) Inject(target []string) ([]string, error) {
	return forceQuotesReplace(target, dsSummaryInjectorTarget,
		d.state.Summary), nil
}
