package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDateInjector_Inject(t *testing.T) {
	Convey("Date Injector", t, func() {
		tmp, _ := time.Parse(time.RFC3339, "1988-10-31T04:38:13Z")
		testDate := "1988-10-31"

		details := job.Details{
			StorableDetails: job.StorableDetails{Started: tmp.UTC()},
		}

		tests := [][2][]string{
			{{"<<date>>"}, {testDate}},
			{{`"<<date>>"`}, {`"` + testDate + `"`}},
			{{"--foo=<<date>>"}, {`--foo=` + testDate}},
			{{`--foo="<<date>>"`}, {`--foo="` + testDate + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDateInjector(&details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
