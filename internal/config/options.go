package config

import (
	"strconv"

	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

const (
	OptKeyServiceNameYaml = "service-name"
	OptKeyCommandsYaml    = "commands"
)

// Options is a container for the full configuration of the running server.
//
// This includes both CLI params and configuration file contents.
type Options struct {
	ServiceName string  `yaml:"service-name" json:"serviceName"`
	Port        uint16  `yaml:"-" json:"port"`
	ConfigPath  string  `yaml:"-" json:"configPath"`
	Command     Command `yaml:"command" json:"command"`
	Version     string  `yaml:"-" json:"-"`
	Workspace   string  `yaml:"-" json:"workspace"`
}

// GetUsablePort returns the configured server port in the format expected by
// the Golang HTTP server package.
func (o *Options) GetUsablePort() string {
	return ":" + strconv.FormatUint(uint64(o.Port), 10)
}

// Validate confirms that the Options instance contains all the required config
// values needed to start up the server.
func (o *Options) Validate() {
	L := log.Logger()
	errored := false

	if len(o.ServiceName) == 0 {
		L.Error("Config: serviceName is required.")

		errored = true
	}

	if len(o.Command.Executable) == 0 {
		L.Error("Config: at least one command must be configured.")

		errored = true
	}

	err := o.Command.Validate()
	if err != nil {
		L.Errorf("Config: Command: %s", err.Error())

		errored = true
	}

	if errored {
		L.Fatal("Shutting down due to configuration errors.")
	}
}
