package inject

import (
	"fmt"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	dsUserEmailInjectorTarget = "<<ds-user-email>>"
	dsUserEmailFormat         = "handler.%d@veupathdb.org"
)

// NewDsUserEmailInjector returns a new VariableInjector instance that will
// replace <<ds-user-email>> variables in a command config.
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
