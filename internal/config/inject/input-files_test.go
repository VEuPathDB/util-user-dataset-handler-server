package inject_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

func TestInputFileInjector_Inject(t *testing.T) {
	Convey("InputFile Injector", t, func() {
		Convey("Valid Template Variables", func() {
			input1 := "foo.txt"
			input2 := "bar.txt"

			details := job.Details{InputFiles: []string{input1, input2}}

			tests := [][2][]string{
				{{"foo", "<<input-files>>"}, {"foo", input1, input2}},
				{{`"<<input-files>>"`}, {`"` + input1 + " " + input2 + `"`}},
				{{"--foo=<<input-files>>"}, {`--foo=`, input1, input2}},
				{{"--foo=<<input-files>>bar"}, {`--foo=`, input1, input2, "bar"}},
				{{`--foo="<<input-files>>"`}, {`--foo="` + input1 + " " + input2 + `"`}},
				{{"--foo=<<input-files[0]>>"}, {`--foo=` + input1}},
				{{`--foo="<<input-files[0]>>"`}, {`--foo="` + input1 + `"`}},
			}

			for _, test := range tests {
				inj := inject.NewInputFileInjector(&details, nil)
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
		Convey("Invalid Template Variables", func() {
			input1 := "foo.txt"
			input2 := "bar.txt"

			details := job.Details{InputFiles: []string{input1, input2}}

			tests := [][]string{
				{"--foo=<<input-files[2]>>"},
				{`--foo="<<input-files[2]>>"`},
				{`--foo="<<input-files[]>>"`},
				{`--foo="<<input-files[a]>>"`},
			}

			for _, test := range tests {
				inj := inject.NewInputFileInjector(&details, nil)
				a, b := inj.Inject(test)
				So(b, ShouldNotBeNil)
				So(a, ShouldBeNil)
			}
		})
	})
}
