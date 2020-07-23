package inject_test

import (
	"github.com/sirupsen/logrus"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsNameInjector_Inject(t *testing.T) {
	Convey("Dataset Name Injector", t, func() {
		details := job.Metadata{
			Name: "Foo bar fizz buzz",
		}
		tests := [][2][]string{
			{{"<<ds-name>>"}, {`"` + details.Name + `"`}},
			{{`"<<ds-name>>"`}, {`"` + details.Name + `"`}},
			{{"--foo=<<ds-name>>"}, {`--foo="` + details.Name + `"`}},
			{{`--foo="<<ds-name>>"`}, {`--foo="` + details.Name + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDsNameInjector(nil, &details, logrus.WithField("test", true))
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
