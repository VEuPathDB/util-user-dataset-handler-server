package parse

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

// ConfigFile parses the config file set by the CLI options.
func ConfigFile(opts *config.Options) {
	raw, err := ioutil.ReadFile(opts.ConfigPath)
	if err != nil {
		log.Logger().Fatal(err)
	}

	tmp := new(config.Options)
	if err = yaml.Unmarshal(raw, tmp); err != nil {
		log.Logger().Fatal(err)
	}

	opts.Command = tmp.Command
	opts.ServiceName = tmp.ServiceName
}
