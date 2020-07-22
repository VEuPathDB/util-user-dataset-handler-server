package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestDsUserIdInjector_Inject(t *testing.T) {
	Convey("Dataset User ID Injector", t, func() {
		Convey("Expands the user ID template variable to the provided user ID", func() {
			expect := "1234"
			details := job.Metadata{BaseInfo: dataset.BaseInfo{Owner: 1234}}

			tests := [][2][]string{
				{{"<<ds-user-id>>"}, {expect}},
				{{`"<<ds-user-id>>"`}, {`"` + expect + `"`}},
				{{"--foo=<<ds-user-id>>"}, {`--foo=` + expect}},
				{{`--foo="<<ds-user-id>>"`}, {`--foo="` + expect + `"`}},
			}

			for _, test := range tests {
				inj := inject.NewDsUserIdInjector(nil, &details)
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
	})
}
