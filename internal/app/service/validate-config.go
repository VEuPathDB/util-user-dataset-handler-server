package service

import (
	"fmt"
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
	errConfNoComs = `Key "` + config.OptKeyCommandsYaml +
		`" must exist and contain at least one valid command entry.`
	errConfComNoPath = `Key "` + config.OptKeyCommandsYaml + `[%d].` +
		config.CmdKeyCommandYaml + `" must exist and be a non empty value.`
	noteConfComNoArgs = `Key "` + config.OptKeyCommandsYaml + `[%d].` +
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

	if len(opts.Commands) == 0 {
		L.Error(errConfNoComs)
		ok = false
	}

	for i := range opts.Commands {
		if !validateCommand(L, &opts.Commands[i], i) {
			ok = false
		}
	}

	if !ok {
		L.Error(errBadConfig)
		os.Exit(app.StatusValidateConfFailed)
	} else {
		L.Info(noteOkConfig)
		os.Exit(app.StatusSuccess)
	}
}

func validateCommand(log *logrus.Entry, cmd *config.Command, index int) bool {
	ok := true
	if len(cmd.Command) == 0 {
		log.Error(fmt.Sprintf(errConfComNoPath, index))
		ok = false
	}

	if len(cmd.Args) == 0 {
		log.Info(fmt.Sprintf(noteConfComNoArgs, index))
	}
	return ok
}
