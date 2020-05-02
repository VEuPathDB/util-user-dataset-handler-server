package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimeInjector_Inject(t *testing.T) {
	Convey("Time Injector", t, func() {
		tmp, _ := time.Parse(time.RFC3339, "1988-10-31T04:38:13Z")
		testTime := "04:38:13"

		details := process.Details{
			StorableDetails: process.StorableDetails{Started: tmp.UTC()},
		}

		tests := [][2][]string{
			{{"<<time>>"}, {testTime}},
			{{`"<<time>>"`}, {`"` + testTime + `"`}},
			{{"--foo=<<time>>"}, {`--foo=` + testTime}},
			{{`--foo="<<time>>"`}, {`--foo="` + testTime + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewTimeInjector(&details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
