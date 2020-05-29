package config

import (
	"errors"
	"fmt"
	"os/exec"
)

const (
	errNoCmd       = "command name is required"
	errCmdNotAvail = "cannot find command with the name %s: %s"
)

// Yaml keys for Command props.
const (
	CmdKeyCommandYaml = "command"
	CmdKeyArgsYaml    = "args"
)

type Command struct {
	Executable string   `yaml:"executable" json:"executable"`
	Args       []string `yaml:"args" json:"arguments"`
}

func (c *Command) Validate() error {
	if len(c.Executable) == 0 {
		return errors.New(errNoCmd)
	}

	if _, err := exec.LookPath(c.Executable); err != nil {
		return fmt.Errorf(errCmdNotAvail, c.Executable, err)
	}

	return nil
}
