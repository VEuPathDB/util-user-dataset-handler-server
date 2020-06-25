package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"strconv"
)

const dsUserIdInjectorTarget = "<<ds-user-id>>"

// NewDsUserIdInjector returns a new VariableInjector instance that will
// replace <<ds-user-email>> variables in a command config.
func NewDsUserIdInjector(_ *job.Details, meta *job.Metadata) VariableInjector {
	return &dsUserIdInjector{meta}
}

type dsUserIdInjector struct {
	state *job.Metadata
}

func (d *dsUserIdInjector) Inject(target []string) ([]string, error) {
	return simpleReplace(target, dsUserIdInjectorTarget,
		strconv.FormatUint(uint64(d.state.Owner), 10)), nil
}
