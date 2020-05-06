package parse

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"io/ioutil"

	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
)

func ConfigFile(opts *config.Options) {
	raw, err := ioutil.ReadFile(opts.ConfigPath)
	if err != nil {
		log.Logger().Fatal(err)
	}

	tmp := new(config.Options)
	err = yaml.Unmarshal(raw, tmp)
	if err != nil {
		log.Logger().Fatal(err)
	}

	opts.Commands = tmp.Commands
	opts.ServiceName = tmp.ServiceName
}
