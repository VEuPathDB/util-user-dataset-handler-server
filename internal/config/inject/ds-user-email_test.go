package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsUserEmailInjector_Inject(t *testing.T) {
	Convey("Dataset UserEmail Injector", t, func() {
		expect := "handler.1234@veupathdb.org"
		details := job.Metadata{
			BaseInfo: dataset.BaseInfo{
				Owner: 1234,
			},
		}
		tests := [][2][]string{
			{{"<<ds-user-email>>"}, {expect}},
			{{`"<<ds-user-email>>"`}, {`"` + expect + `"`}},
			{{"--foo=<<ds-user-email>>"}, {`--foo=` + expect}},
			{{`--foo="<<ds-user-email>>"`}, {`--foo="` + expect + `"`}},
		}

		for _, test := range tests {
			inj := inject.NewDsUserEmailInjector(nil, &details)
			a, b := inj.Inject(test[0])
			So(b, ShouldBeNil)
			So(a, ShouldResemble, test[1])
		}
	})
}
