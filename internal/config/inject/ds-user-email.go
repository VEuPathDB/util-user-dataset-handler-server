package inject

import (
	"fmt"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	dsUserEmailInjectorTarget = "<<ds-user-email>>"
	dsUserEmailFormat         = "handler.%d@veupathdb.org"
)

func NewDsUserEmailInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsUserEmailInjector{meta}
}

type dsUserEmailInjector struct {
	state *job.Metadata
}

func (d *dsUserEmailInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsUserEmailInjectorTarget,
		fmt.Sprintf(dsUserEmailFormat, d.state.Owner)), nil
}
