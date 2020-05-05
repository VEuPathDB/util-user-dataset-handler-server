package util

import (
	"os/exec"

	"github.com/sirupsen/logrus"

	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

// PrepCommand prepares a command to run using the server's
// logging mechanism for it's own stdout and stderr.
func PrepCommand(logger *logrus.Entry, com string, args... string) *exec.Cmd {
	sOut, sErr := log.StdWriters(logger, com)
	cmd := exec.Command(com, args...)
	cmd.Stdout = sOut
	cmd.Stderr = sErr
	return cmd
}
