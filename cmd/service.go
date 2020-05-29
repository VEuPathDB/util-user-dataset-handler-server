package main

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/config/parse"
	"github.com/VEuPathDB/util-exporter-server/internal/server"
)

var version = "untagged dev build"

func main() {
	options := new(config.Options)
	options.Version = version

	log.SetLogger(log.ConfigureLogger())
	parse.Cli(options)
	parse.ConfigFile(options)

	options.Validate()

	log.SetLogger(log.ConfigureLogger().WithField("source", options.ServiceName))

	serve := server.NewServer(options)

	serve.RegisterEndpoints()
	log.Logger().Fatal(serve.Serve())
}
