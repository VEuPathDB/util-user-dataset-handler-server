package log

import (
	"github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
)

func ConfigureLogger() *logrus.Entry {
	log := logrus.New()
	fmt := new(prefixed.TextFormatter)
	fmt.FullTimestamp = true
	fmt.TimestampFormat = "2006-01-02 15:04:05.000000"
	log.Formatter = fmt
	log.Level = logrus.TraceLevel

	return log.WithFields(logrus.Fields{
		FieldSource: "server",
	})
}
