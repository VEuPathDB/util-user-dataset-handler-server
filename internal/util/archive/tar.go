package archive

import (
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/sirupsen/logrus"
	"strings"
)

const (
	CmdTar = "tar"
)

func TarExtensions() []string {
	return []string{
		".tar",
		".tar.gz",
		".tgz",
	}
}

func HasTarExtension(file string) bool {
	file = strings.ToLower(file)

	for _, ext := range TarExtensions() {
		if strings.HasSuffix(file, ext) {
			return true
		}
	}

	return false
}

func UnTar(dir, file string, log *logrus.Entry) ([]string, error) {
	return unpack(dir, file, util.PrepCommand(log, CmdTar, "-xf", file))
}
