package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
)

const (
	indentSize    = 2
	headerComment = `# Config Template
#
# Please edit the variables below and rename this file.

`
)

func main() {
	L := log.ConfigureLogger().WithField(log.FieldSource, "gen-config")

	file, err := os.Create("config.tpl.yml")
	checkErr(L, err)
	defer file.Close()

	_, err = file.Write([]byte(headerComment))
	checkErr(L, err)

	enc := yaml.NewEncoder(file)
	enc.SetIndent(indentSize)

	checkErr(L, enc.Encode(config.Options{
		ServiceName: "my service",
		FileTypes: []string{".txt"},
		Command: config.Command{
			Executable: "my-command",
			Args:       []string{"<<input-files>>"},
		},
	}))
}

func checkErr(log *logrus.Entry, err error) {
	if err != nil {
		log.Fatal(err)
	}
}