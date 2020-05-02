package inject_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/config/inject"
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestOutputFileInjector_Inject(t *testing.T) {
	Convey("OutputFile Injector", t, func() {
		input1 := "foo.txt"
		input2 := "bar.txt"
		input3 := "fizz.txt"
		input4 := "buzz.txt"

		details := process.Details{
			OutputFiles: [][]string{{input1, input2}, {input3, input4}},
		}

		Convey("Valid Template Variables", func() {
			tests := [][2][]string{
				{{"foo", "<<output-files>>"}, {"foo", input3, input4}},
				{{`"<<output-files>>"`}, {`"` + input3 + " " + input4 + `"`}},
				{{"--foo=<<output-files>>"}, {`--foo=`, input3, input4}},
				{{"--foo=<<output-files>>bar"}, {`--foo=`, input3, input4, "bar"}},
				{{`--foo="<<output-files>>"`}, {`--foo="` + input3 + " " + input4 + `"`}},
				{{"--foo=<<output-files[0]>>"}, {`--foo=`, input1, input2}},
				{{"--foo=<<output-files[0]>>bar"}, {`--foo=`, input1, input2, "bar"}},
				{{`--foo="<<output-files[0]>>"`}, {`--foo="` + input1 + " " + input2 + `"`}},
				{{"--foo=<<output-files[0][1]>>"}, {`--foo=` + input2}},
				{{`--foo="<<output-files[0][1]>>"`}, {`--foo="` + input2 + `"`}},
			}

			for _, test := range tests {
				inj := inject.NewOutputFileInjector(&details)
				a, b := inj.Inject(test[0])
				So(b, ShouldBeNil)
				So(a, ShouldResemble, test[1])
			}
		})
		Convey("Invalid Template Variables", func() {
			tests := [][]string{
				{"--foo=<<output-files[2]>>"},
				{`--foo="<<output-files[2]>>"`},
				{`--foo="<<output-files[]>>"`},
				{`--foo="<<output-files[a]>>"`},
				{"--foo=<<output-files[0][2]>>"},
				{`--foo="<<output-files[0][2]>>"`},
				{"--foo=<<output-files[2][2]>>"},
				{`--foo="<<output-files[2][2]>>"`},
				{`--foo="<<output-files[0][]>>"`},
				{`--foo="<<output-files[0][a]>>"`},
				{`--foo="<<output-files[a][0]>>"`},
			}

			for _, test := range tests {
				inj := inject.NewOutputFileInjector(&details)
				a, b := inj.Inject(test)
				So(b, ShouldNotBeNil)
				So(a, ShouldBeNil)
			}
		})
	})
}
