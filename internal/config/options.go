package config

import (
	"github.com/sirupsen/logrus"
)

const (
	OptKeyServiceNameYaml = "service-name"
	OptKeyCommandsYaml    = "commands"
)

// Options is a container for the full configuration of the running server.
//
// This includes both CLI params and configuration file contents.
type Options struct {

	// Name of the service as it appears in health/status output.
	ServiceName string   `yaml:"service-name" json:"serviceName"`

	// Configuration for the CLI call to the dataset tooling used to process
	// uploads.
	Command     Command  `yaml:"command" json:"command"`

	// Allowed list of raw upload file extensions.
	//
	// The service will handle zip and tar archives automatically.  This config
	// option simply controls what file extensions are allowed for uploads in
	// addition to the server's built in archive handling.
	//
	// Leaving this empty, or omitting it will mean that the server will only
	// accept it's own known archive formats, and will reject everything else.
	//
	// Setting this to, for example, ".txt" will result in the server allowing
	// file uploads for archive formats _and_ files ending in the extension
	// ".txt".
	FileTypes   []string `yaml:"file-types" json:"fileTypes"`

	// CLI Option: Port the HTTP server will bind to on startup.
	Port        uint16   `yaml:"-" json:"port"`

	// CLI Option: Path to the configuration file to use.
	ConfigPath  string   `yaml:"-" json:"configPath"`

	// Internal: The current service binary version as a string value.
	Version     string   `yaml:"-" json:"-"`

	Workspace   string   `yaml:"-" json:"workspace"`
}


// Validate confirms that the Options instance contains all the required config
// values needed to start up the server.
func IsValid(log *logrus.Entry, o FileOptions) (errored bool) {
	errored = true

	if len(o.ServiceName()) == 0 {
		log.Error("Config: serviceName is required.")

		errored = false
	}

	if len(o.Commands().Executable) == 0 {
		log.Error("Config: at least one command must be configured.")

		errored = false
	}

	err := o.Commands().Validate()
	if err != nil {
		log.Errorf("Config: Command: %s", err.Error())

		errored = false
	}

	return
}
