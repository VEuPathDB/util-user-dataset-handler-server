package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsSummaryInjectorTarget = "<<ds-summary>>"

func NewDsSummaryInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsSummaryInjector{meta}
}

type dsSummaryInjector struct {
	state *job.Metadata
}

func (d *dsSummaryInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsSummaryInjectorTarget,
		d.state.Summary), nil
}
