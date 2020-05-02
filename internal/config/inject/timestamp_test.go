package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTimestampInjector_Inject(t *testing.T) {
	Convey("Timestamp Injector", t, func() {
		tmp, _ := time.Parse(time.RFC3339, "1988-10-31T04:38:13Z")
		testTimestamp := "594275893"

		details := process.Details{
			StorableDetails: process.StorableDetails{Started: tmp.UTC()},
		}

		tests := [][2][]string{
			{{"<<timestamp>>"}, {testTimestamp}},
			{{`"<<timestamp>>"`}, {`"` + testTimestamp + `"`}},
			{{"--foo=<<timestamp>>"}, {`--foo=` + testTimestamp}},
			{{`--foo="<<timestamp>>"`}, {`--foo="` + testTimestamp + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewTimestampInjector(&details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
