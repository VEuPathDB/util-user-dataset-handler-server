package config_test

import (
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
	"github.com/sirupsen/logrus"
	"go/build"
	"os"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
)

// TODO: Improve this test, this dangerously assumes that files and injectors
//       are 1 to 1.
func TestInjectorList(t *testing.T) {
	// Files that should not be counted as usable Injector implementations
	miscFiles := []string{
		"injector.go",
		"util.go",
	}

	Convey("Injector Listing", t, func() {
		cwd, err := os.Getwd()
		So(err, ShouldBeNil)

		pkg, err := build.Import(
			"github.com/VEuPathDB/util-exporter-server/internal/config/inject",
			cwd, build.IgnoreVendor)
		So(err, ShouldBeNil)

		files := make([]string, 0, len(pkg.GoFiles)-len(miscFiles))
	outer:
		for _, file := range pkg.GoFiles {
			for _, exc := range miscFiles {
				if file == exc {
					continue outer
				}
			}
			files = append(files, file)
		}

		So(len(config.InjectorList()), ShouldEqual, len(files))
	})
}

func TestFullInjectionBiom(t *testing.T) {
	now := time.Now()

	testDetails := job.Details{
		StorableDetails: job.StorableDetails{
			Started:  &now,
			UserID:   1234,
			Token:    "abcdabcdabcd",
			Status:   job.StatusProcessing,
			Projects: []site.WdkSite{site.MicrobiomeDB},
		},
		InputFile:     "fruit.tar",
		UnpackedFiles: []string{"apple", "banana", "orange"},
		WorkingDir:    "/workspace/potatoes",
	}

	testMeta := job.Metadata{
		BaseInfo: dataset.BaseInfo{
			Projects: []site.WdkSite{site.MicrobiomeDB},
			Owner:    1234,
		},
		Name:        "strawberry pancakes",
		Summary:     "blueberry pancakes",
		Description: "blackberry pancakes",
		Token:       "abcdabcdabcd",
		Origin:      "mars",
	}

	configParams := []string{
		"<<ds-name>>",
		"<<ds-summary>>",
		"<<ds-description>>",
		`"<<ds-user-id>>"`,
		"html.html",
		`"<<ds-origin>>"`,
		"<<input-files[0]>>",
	}

	Convey("Full config test", t, func() {
		var err error
		out := configParams
		log := logrus.WithField("test", true)
		for _, fn := range config.InjectorList() {
			tmp := fn(&testDetails, &testMeta, log)
			out, err = tmp.Inject(out)
		}

		So(err, ShouldBeNil)
		So(out, ShouldResemble, []string{
			`"` + testMeta.Name + `"`,
			`"` + testMeta.Summary + `"`,
			`"` + testMeta.Description + `"`,
			`"1234"`,
			"html.html",
			`"mars"`,
			"apple",
		})
	})
}
