package config_test

import (
	"go/build"
	"os"
	"testing"

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
