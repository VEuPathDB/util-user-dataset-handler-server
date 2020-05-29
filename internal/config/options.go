package config

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"strconv"
)

const (
	OptKeyServiceNameYaml = "service-name"
	OptKeyCommandsYaml    = "commands"
)

type Options struct {
	ServiceName string  `yaml:"service-name" json:"serviceName"`
	Port        uint16  `yaml:"-" json:"port"`
	ConfigPath  string  `yaml:"-" json:"configPath"`
	Command     Command `yaml:"command" json:"command"`
	Version     string  `yaml:"-" json:"-"`
	Workspace   string  `yaml:"-" json:"workspace"`
}

func (O *Options) GetUsablePort() string {
	return ":" + strconv.FormatUint(uint64(O.Port), 10)
}

func (O *Options) Validate() {
	L := log.Logger()
	errored := false
	if len(O.ServiceName) == 0 {
		L.Error("Config: serviceName is required.")
		errored = true
	}

	if len(O.Command.Executable) == 0 {
		L.Error("Config: at least one command must be configured.")
		errored = true
	}

	err := O.Command.Validate()
	if err != nil {
		L.Errorf("Config: Command: %s", err.Error())
		errored = true
	}

	if errored {
		L.Fatal("Shutting down due to configuration errors.")
	}
}
