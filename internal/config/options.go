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
	ServiceName string    `yaml:"service-name" json:"serviceName"`
	Port        uint16    `yaml:"-" json:"port"`
	ConfigPath  string    `yaml:"-" json:"configPath"`
	Commands    []Command `yaml:"commands" json:"commands"`
	Version     string    `yaml:"-" json:"-"`
	Workspace   string    `yaml:"-" json:"workspace"`
}

func (O *Options) GetUsablePort() string {
	return ":" + strconv.FormatUint(uint64(O.Port), 10)
}

func (O *Options) Validate() {
	L := log.ConfigureLogger("service", "booting")
	errored := false
	if len(O.ServiceName) == 0 {
		L.Error("Config: serviceName is required.")
		errored = true
	}

	if len(O.Commands) == 0 {
		L.Error("Config: at least one command must be configured.")
		errored = true
	}

	for i := range O.Commands {
		err := O.Commands[i].Validate()
		if err != nil {
			L.Errorf("Config: Command %d: %s", i, err.Error())
			errored = true
		}
	}

	if errored {
		L.Fatal("Shutting down due to configuration errors.")
	}
}
