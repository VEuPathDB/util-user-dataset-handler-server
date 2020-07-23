package inject_test

import (
	"github.com/sirupsen/logrus"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsSummaryInjector_Inject(t *testing.T) {
	Convey("Dataset Summary Injector", t, func() {
		details := job.Metadata{
			Summary: "Foo bar fizz buzz",
		}
		tests := [][2][]string{
			{{"<<ds-summary>>"}, {`"` + details.Summary + `"`}},
			{{`"<<ds-summary>>"`}, {`"` + details.Summary + `"`}},
			{{"--foo=<<ds-summary>>"}, {`--foo="` + details.Summary + `"`}},
			{{`--foo="<<ds-summary>>"`}, {`--foo="` + details.Summary + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDsSummaryInjector(nil, &details, logrus.WithField("test", true))
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
