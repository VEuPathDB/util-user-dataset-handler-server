package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const dsDescriptionInjectorTarget = "<<ds-description>>"

func NewDsDescriptionInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsDescriptionInjector{meta}
}

type dsDescriptionInjector struct {
	state *job.Metadata
}

func (d *dsDescriptionInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsDescriptionInjectorTarget,
		d.state.Description), nil
}
