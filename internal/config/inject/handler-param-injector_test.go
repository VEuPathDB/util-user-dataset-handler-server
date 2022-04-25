package inject_test

import (
	"testing"

	"github.com/sirupsen/logrus"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestHandlerParamInjector_Inject(t *testing.T) {
	Convey("HandlerParam Injector", t, func() {
		Convey("Expands valid template variables", func() {
			metadata := job.Metadata{HandlerParams: map[string]string{
				"p1": "param1",
				"p2": "param2"}}

			tests := [][2][]string{
				{{"foo", "<<handler-params.p1>>"}, {"foo", "param1"}},
				{{`"<<handler-params.p2>>"`}, {"param2"}},
				{{"<<handler-params.p2.p3>>"}, {"<<handler-params.p2.p3>>"}}, // No nesting of params supported by injector, error here?
				{{"--foo=<<handler-params.p1>>"}, {"--foo=param1"}},
				{{"--foo=<<handler-params.p2>>bar"}, {"--foo=param2bar"}},
				{{"--foo=<<handler-params>>"}, {"--foo=<<handler-params>>"}}, // No substitution made unless "." is present, error here?
			}

			for _, test := range tests {
				inj := inject.NewHandlerParamInjector(nil, &metadata, logrus.WithField("test", true))
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
		Convey("Errors on invalid template variables", func() {
			metadata := job.Metadata{HandlerParams: map[string]string{
				"p1": "param1",
				"p2": "param2"}}

			tests := [][]string{
				{"--foo=<<handler-params.unknown>>"},
			}

			for _, test := range tests {
				inj := inject.NewHandlerParamInjector(nil, &metadata, logrus.WithField("test", true))
				a, b := inj.Inject(test)
				So(a, ShouldBeNil)
				So(b, ShouldNotBeNil)
			}
		})
		Convey("Errors when params missing from job metadata", func() {
			metadata := job.Metadata{}

			tests := [][]string{
				{"--foo=<<handler-params.p1>>"},
			}

			for _, test := range tests {
				inj := inject.NewHandlerParamInjector(nil, &metadata, logrus.WithField("test", true))
				a, b := inj.Inject(test)
				So(a, ShouldBeNil)
				So(b, ShouldNotBeNil)
			}
		})
	})
}
