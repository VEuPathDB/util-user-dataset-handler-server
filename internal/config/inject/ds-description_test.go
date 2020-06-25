package inject_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsDescriptionInjector_Inject(t *testing.T) {
	Convey("Dataset Description Injector", t, func() {
		details := job.Metadata{
			Description: "Foo bar fizz buzz",
		}
		tests := [][2][]string{
			{{"<<ds-description>>"}, {`"` + details.Description + `"`}},
			{{`"<<ds-description>>"`}, {`"` + details.Description + `"`}},
			{{"--foo=<<ds-description>>"}, {`--foo="` + details.Description + `"`}},
			{{`--foo="<<ds-description>>"`}, {`--foo="` + details.Description + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDsDescriptionInjector(nil, &details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
