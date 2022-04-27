package inject_test

import (
	"testing"

	"github.com/sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
)

func TestProjectsInjector_Test(t *testing.T) {
	Convey("Projects Injector", t, func() {
		Convey("Single Project", func() {
			projects := []site.WdkSite{site.VectorBase}
			baseInfo := dataset.BaseInfo{Projects: projects}
			metadata := job.Metadata{BaseInfo: baseInfo}

			tests := [][2][]string{
				{{"<<projects>>"}, {"VectorBase"}},
				{{"--projects=<<projects>>"}, {"--projects=VectorBase"}},
			}

			for _, test := range tests {
				inj := inject.NewProjectsInjector(nil, &metadata, logrus.WithField("test", true))
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
		Convey("Multiple Projects", func() {
			projects := []site.WdkSite{site.VectorBase, site.PlasmoDB}
			baseInfo := dataset.BaseInfo{Projects: projects}
			metadata := job.Metadata{BaseInfo: baseInfo}

			tests := [][2][]string{
				{{"<<projects>>"}, {"VectorBase,PlasmoDB"}},
				{{"--projects=<<projects>>"}, {"--projects=VectorBase,PlasmoDB"}},
			}

			for _, test := range tests {
				inj := inject.NewProjectsInjector(nil, &metadata, logrus.WithField("test", true))
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
	})
}
