package archive

import (
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	CmdUnzip = "unzip"
)

func ZipExtensions() []string {
	return []string{
		".zip",
	}
}

func HasZipExtension(file string) bool {
	file = strings.ToLower(file)

	for _, ext := range ZipExtensions() {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}

	return false
}

func UnZip(dir, file string, log *logrus.Entry) ([]string, error) {
	return unpack(dir, file, util.PrepCommand(log, CmdUnzip, "-xf", file))
}