package log

import (
	"github.com/sirupsen/logrus"
	"io"
)

// StdWriters constructs a pair of log writers that pass
// command output through to the standard logger.
func StdWriters(log *logrus.Entry, cmd string) (out, err io.Writer) {
	log = log.WithField("command", cmd)
	return &stdWriter{fn: log.Debug}, &stdWriter{fn: log.Error}
}

type stdWriter struct {
	fn  func(...interface{})
	buf []byte
}

func (s *stdWriter) Write(p []byte) (n int, err error) {
	s.fn(string(p))
	return len(p), nil
}


