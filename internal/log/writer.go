package log

import (
	"io"

	"github.com/sirupsen/logrus"
)

// StdWriters constructs a pair of log writers that pass
// command output through to the standard logger.
func StdWriters(log *logrus.Entry, cmd string) (out, err io.Writer) {
	log = log.WithField("source", cmd)
	return &stdWriter{fn: log.Debug}, &stdWriter{fn: log.Error}
}

type stdWriter struct {
	fn func(...interface{})
}

func (s *stdWriter) Write(p []byte) (n int, err error) {
	s.fn(string(p))
	return len(p), nil
}
