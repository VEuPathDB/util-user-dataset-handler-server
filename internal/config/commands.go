package config

import (
	"errors"
	"fmt"
	"os/exec"
)

const (
	errNoCmd       = "Command name is required"
	errCmdNotAvail = "Cannot find command with the name %s: %s"
)

// Yaml keys for Command props
const (
	CmdKeyCommandYaml = "command"
	CmdKeyArgsYaml    = "args"
)

type Command struct {
	Executable string   `yaml:"executable" json:"executable"`
	Args       []string `yaml:"args" json:"arguments"`
}

func (C *Command) Validate() error {
	if len(C.Executable) == 0 {
		return errors.New(errNoCmd)
	}

	if _, err := exec.LookPath(C.Executable); err != nil {
		return fmt.Errorf(errCmdNotAvail, C.Executable, err)
	}

	return nil
}
