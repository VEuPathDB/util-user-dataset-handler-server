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
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args" json:"arguments"`
}

func (C *Command) Validate() error {
	if len(C.Command) == 0 {
		return errors.New(errNoCmd)
	}

	_, err := exec.LookPath(C.Command)
	if err != nil {
		return fmt.Errorf(errCmdNotAvail, C.Command, err)
	}

	return nil
}
