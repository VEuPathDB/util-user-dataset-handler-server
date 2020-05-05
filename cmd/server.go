package main

import (
	"github.com/sirupsen/logrus"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

var version = "untagged dev build"

func main() {
	options := new(config.Options)
	options.Version = version
	config.ParseCli(options)
	config.ParseOptions(options)

	prepareLogger(options)

	serve := server.NewServer(options)
	serve.RegisterEndpoints()
	util.Logger().Fatal(serve.Serve())
}

func prepareLogger(opts *config.Options) {
	logrus.SetFormatter(new(logrus.TextFormatter))
	util.SetLogger(logrus.StandardLogger().
		WithField("service", opts.ServiceName))
}
