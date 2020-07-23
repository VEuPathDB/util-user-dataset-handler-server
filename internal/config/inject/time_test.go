package inject_test

import (
	"github.com/sirupsen/logrus"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestTimeInjector_Inject(t *testing.T) {
	Convey("Time Injector", t, func() {
		tmp, _ := time.Parse(time.RFC3339, "1988-10-31T04:38:13Z")
		testTime := "04:38:13"
		utc := tmp.UTC()

		details := job.Details{
			StorableDetails: job.StorableDetails{Started: &utc},
		}

		tests := [][2][]string{
			{{"<<time>>"}, {testTime}},
			{{`"<<time>>"`}, {`"` + testTime + `"`}},
			{{"--foo=<<time>>"}, {`--foo=` + testTime}},
			{{`--foo="<<time>>"`}, {`--foo="` + testTime + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewTimeInjector(&details, nil, logrus.WithField("test", true))
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
