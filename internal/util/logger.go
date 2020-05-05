package util

import "github.com/sirupsen/logrus"

var logger *logrus.Entry

func SetLogger(entry *logrus.Entry) {
	logger = entry
}

func Logger() *logrus.Entry {
	return logger
}
