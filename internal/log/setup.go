package log

import (
	"github.com/VEuPathDB/util-exporter-server/internal/app"
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func ConfigureLogger(svcName, status string) *logrus.Entry {
	log := logrus.New()
	log.Formatter = new(prefixed.TextFormatter)
	return log.WithFields(logrus.Fields{
		app.Keys.Logger.Service: svcName,
		app.Keys.Logger.Status:  status,
	})
}
