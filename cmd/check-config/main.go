package main

import (
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

const (
	errConfSvcName     = `Key "` + config.OptKeyServiceNameYaml +
		`" must exist and be a non-empty value.`
	errConfComNoPath = `Key "` + config.OptKeyCommandsYaml +
		config.CmdKeyCommandYaml + `" must exist and be a non empty value.`
	noteConfComNoArgs = `Key "` + config.OptKeyCommandsYaml +
		config.CmdKeyArgsYaml + `" has no arguments.  It will receive no input ` +
		`from the server at runtime.`
	errBadConfig = "Config invalid."
	noteOkConfig = "Config valid."
	noteInfo     = `Validating the configuration file "%s".  NOTE: This ` +
		"validation does not verify that the configured commands exist on the " +
		"current $PATH."
)

func main() {
	L := log.ConfigureLogger().WithField(log.FieldSource, "check-config")

	cliOpts, err := config.ParseCLIOptions()
	checkErr(L, err)

	L.Infof(noteInfo, cliOpts.ConfigPath())

	bytes, err := ioutil.ReadFile(cliOpts.ConfigPath())
	checkErr(L, err)

	opts := new(config.Options)
	checkErr(L, yaml.Unmarshal(bytes, opts))

	ok := true

	if opts.ServiceName == "" {
		L.Error(errConfSvcName)
		ok = false
	}

	if !validateCommand(L, &opts.Command) {
		ok = false
	}

	if !ok {
		L.Fatal(errBadConfig)
	}

	L.Info(noteOkConfig)
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

func checkErr(log *logrus.Entry, err error) {
	if err != nil {
		log.Fatal(err)
	}
}
