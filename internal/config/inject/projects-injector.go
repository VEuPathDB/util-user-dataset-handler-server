package inject

import (
	"strings"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
	"github.com/sirupsen/logrus"
)

const projectsInjectorTarget = "<<projects>>"

// NewDsUserIdInjector returns a new VariableInjector instance that will
// replace <<projects>> variables in a command config.
func NewProjectsInjector(
	_ *job.Details,
	meta *job.Metadata,
	log *logrus.Entry,
) VariableInjector {
	log.Trace("inject.NewProjectsInjector")
	return &projectsInjector{log, meta}
}

type projectsInjector struct {
	log   *logrus.Entry
	state *job.Metadata
}

func (injector *projectsInjector) Inject(target []string) ([]string, error) {
	injector.log.Trace("inject.projectsInjector.Inject")
	return simpleReplace(target, projectsInjectorTarget,
		strings.Join(sitesToString(injector.state.Projects), ",")), nil
}

func sitesToString(sites []site.WdkSite) []string {
	stringSites := []string{}
	for _, site := range sites {
		stringSites = append(stringSites, string(site))
	}
	return stringSites
}
