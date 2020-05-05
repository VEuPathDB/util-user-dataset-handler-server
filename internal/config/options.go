package config

import "strconv"

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

type Command struct {
	Command string   `yaml:"command" json:"command"`
	Args    []string `yaml:"args" json:"arguments"`
}