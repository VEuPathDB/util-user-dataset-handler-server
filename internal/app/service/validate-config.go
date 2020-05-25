package service

import (
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/app"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

const (
	ValidateAppName = "check-config"
)

const (
	errReadConfigFail  = "Failed to read config file.  Process said:"
	errParseConfigFail = "Failed to parse config file as Yaml.  Process said:"
	errConfSvcName     = `Key "` + config.OptKeyServiceNameYaml +
		`" must exist and be a non-empty value.`
	errConfComNoPath = `Key "` + config.OptKeyCommandsYaml +
		config.CmdKeyCommandYaml + `" must exist and be a non empty value.`
	noteConfComNoArgs = `Key "` + config.OptKeyCommandsYaml +
		config.CmdKeyArgsYaml + `" has no arguments.  It will receive no input ` +
		`from the server at runtime.`
	errBadConfig = "Config invalid."
	noteOkConfig = "Config valid."
	noteInfo = "Validating the configuration file \"%s\".  NOTE: This " +
		"validation does not verify that the configured commands exist on the" +
		"current $PATH."
)

// ValidateConfig takes an Options object containing just
// the options parsed from the command line and attempts to
// read and validate the Yaml server config file.
func ValidateConfig(options *config.Options) {
	L := log.ConfigureLogger(ValidateAppName, "running")

	L.Info(noteInfo)

	bytes, err := ioutil.ReadFile(options.ConfigPath)
	if err != nil {
		L.Fatal(errReadConfigFail, err)
	}

	opts := new(config.Options)
	err = yaml.Unmarshal(bytes, opts)
	if err != nil {
		L.Fatal(errParseConfigFail, err)
	}

	ok := true
	if opts.ServiceName == "" {
		L.Error(errConfSvcName)
		ok = false
	}

	if !validateCommand(L, &opts.Command) {
		ok = false
	}

	if !ok {
		L.Error(errBadConfig)
		os.Exit(app.StatusValidateConfFailed)
	} else {
		L.Info(noteOkConfig)
		os.Exit(app.StatusSuccess)
	}
}

func validateCommand(log *logrus.Entry, cmd *config.Command) bool {
	ok := true
	if len(cmd.Executable) == 0 {
		log.Error(errConfComNoPath)
		ok = false
	}

	if len(cmd.Args) == 0 {
		log.Info(noteConfComNoArgs)
	}
	return ok
}
