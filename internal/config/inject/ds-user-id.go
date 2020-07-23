package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
	"strconv"
)

const dsUserIdInjectorTarget = "<<ds-user-id>>"

// NewDsUserIdInjector returns a new VariableInjector instance that will
// replace <<ds-user-id>> variables in a command config.
func NewDsUserIdInjector(
	_ *job.Details,
	meta *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewDsUserIdInjector")
	return &dsUserIdInjector{log, meta}
}

type dsUserIdInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (d *dsUserIdInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.dsUserIdInjector.Inject")
	return simpleReplace(target, dsUserIdInjectorTarget,
		strconv.FormatUint(uint64(d.state.Owner), 10)), nil
}
