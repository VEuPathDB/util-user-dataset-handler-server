package main

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server"
)

func main() {
	logger := log.ConfigureLogger()
	log.SetLogger(logger)

	cliOpts, err := config.ParseCLIOptions()
	checkErr(err)

	fileOpts, err := config.ParseFileOptions(cliOpts.ConfigPath())
	checkErr(err)

	if !config.IsValid(logger, fileOpts) {
		logger.Fatal("Shutting down due to configuration errors.")
	}

	log.SetLogger(logger.WithField(log.FieldSource, fileOpts.ServiceName))

	serve := server.NewServer(cliOpts, fileOpts)

	serve.RegisterEndpoints()
	log.Logger().Fatal(serve.Serve())
}

func checkErr(err error) {
	if err != nil {
		log.Logger().Fatal(err)
	}
}
