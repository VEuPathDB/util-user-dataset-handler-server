package inject

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/sirupsen/logrus"
)

const cwdInjectorTarget = "<<cwd>>"

// NewCwdInjector returns a new VariableInjector instance that will replace
// <<cwd>> variables in a command config.
func NewCwdInjector(
	det *job.Details,
	_ *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewCwdInjector")
	return &cwdInjector{log, det}
}

type cwdInjector struct {
	log   *logrus.Entry
	state *job.Details
}

func (d *cwdInjector) Inject(target []string) ([]string, error) {
	d.log.Trace("inject.cwdInjector.Inject")
	return simpleReplace(target, cwdInjectorTarget,
		d.state.WorkingDir), nil
}
