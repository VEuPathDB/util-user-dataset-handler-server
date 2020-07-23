package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const dsSummaryInjectorTarget = "<<ds-summary>>"

// NewDsSummaryInjector returns a new VariableInjector instance that will
// replace <<ds-summary>> variables in a command config.
func NewDsSummaryInjector(
	_ *job.Details,
	meta *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewDsSummaryInjector")
	return &dsSummaryInjector{log, meta}
}

type dsSummaryInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (d *dsSummaryInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dsSummaryInjector.Inject")
	return forceQuotesReplace(target, dsSummaryInjectorTarget,
		d.state.Summary), nil
}
