package service

import (
	"os"

	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/app"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

const (
	GenerateAppName = "gen-config"
)

const (
	indentSize    = 2
	headerComment = `# Config Template
#
# Please edit the variables below and rename this file.

`
)

func GenerateConfig() {
	L := log.ConfigureLogger().WithField("source", GenerateAppName)

	file, err := os.Create("config.tpl.yml")
	if err != nil {
		L.Error(err)
		os.Exit(app.StatusGenConfFailed)
	}
	defer file.Close()

	_, err = file.Write([]byte(headerComment))
	if err != nil {
		L.Error(err)
		os.Exit(app.StatusGenConfFailed)
	}

	enc := yaml.NewEncoder(file)
	enc.SetIndent(indentSize)

	err = enc.Encode(config.Options{
		ServiceName: "my service",
		Command: config.Command{
			Executable: "my-command",
			Args:       []string{"<<input-files>>"},
		},
	})

	if err != nil {
		L.Error(err)
		os.Exit(app.StatusGenConfFailed)
	}

	os.Exit(app.StatusSuccess)
}
