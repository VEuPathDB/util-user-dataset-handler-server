package util

import (
	"os/exec"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

// PrepCommand prepares a command to run using the server's
// logging mechanism for it's own stdout and stderr.
func PrepCommand(logger *logrus.Entry, com string, args ...string) *exec.Cmd {
	sOut, sErr := log.StdWriters(logger, com)
	cmd := exec.Command(com, args...)
	cmd.Stdout = sOut
	cmd.Stderr = sErr

	return cmd
}

// TimeCmd executes the given command and returns the execution time as a float
// value of milliseconds and the response error (if any) from the command
// execution.
func TimeCmd(cmd *exec.Cmd) (millis float64, err error) {
	start := time.Now()
	err = cmd.Run()
	millis = float64(time.Since(start)) / float64(time.Millisecond)

	return
}
