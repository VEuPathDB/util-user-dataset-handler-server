package inject_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsOriginInjector_Inject(t *testing.T) {
	Convey("Dataset Origin Injector", t, func() {
		details := job.Metadata{
			Origin: "galaxy",
		}
		tests := [][2][]string{
			{{"<<ds-origin>>"}, {details.Origin}},
			{{`"<<ds-origin>>"`}, {`"` + details.Origin + `"`}},
			{{"--foo=<<ds-origin>>"}, {`--foo=` + details.Origin}},
			{{`--foo="<<ds-origin>>"`}, {`--foo="` + details.Origin + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDsOriginInjector(nil, &details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
