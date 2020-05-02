package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCwdInjector_Inject(t *testing.T) {
	Convey("Cwd Injector", t, func() {
		testPath := "/test/path"
		details := process.Details{
			WorkingDir: "/test/path",
		}
		tests := [][2][]string{
			{{"<<cwd>>"}, {testPath}},
			{{`"<<cwd>>"`}, {`"` + testPath + `"`}},
			{{"--foo=<<cwd>>"}, {`--foo=` + testPath}},
			{{`--foo="<<cwd>>"`}, {`--foo="` + testPath + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewCwdInjector(&details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
