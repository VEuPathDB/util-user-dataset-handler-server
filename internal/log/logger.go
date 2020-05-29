package log

import "github.com/sirupsen/logrus"

var logger *logrus.Entry

// SetLogger overwrites the global logger base instance with the given value.
func SetLogger(entry *logrus.Entry) {
	logger = entry
}

// Logger returns the global base logger instance.
func Logger() *logrus.Entry {
	return logger
}
