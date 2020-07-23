package inject_test

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDateTimeInjector_Inject(t *testing.T) {
	Convey("DateTime Injector", t, func() {
		tmp, _ := time.Parse(time.RFC3339, "1988-10-31T04:38:13Z")
		testDate := "1988-10-31T04:38:13Z"
		utc := tmp.UTC()

		details := job.Details{
			StorableDetails: job.StorableDetails{Started: &utc},
		}

		tests := [][2][]string{
			{{"<<date-time>>"}, {testDate}},
			{{`"<<date-time>>"`}, {`"` + testDate + `"`}},
			{{"--foo=<<date-time>>"}, {`--foo=` + testDate}},
			{{`--foo="<<date-time>>"`}, {`--foo="` + testDate + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDateTimeInjector(&details, nil, logrus.WithField("test", true))
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
