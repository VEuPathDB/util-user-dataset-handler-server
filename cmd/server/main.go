package main

import (
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server"
)

func main() {
	log.SetLogger(log.ConfigureLogger())

	cliOpts, err := config.ParseCLIOptions()
	checkErr(err)

	fileOpts, err := config.ParseFileOptions(cliOpts.ConfigPath())
	checkErr(err)

	if !config.IsValid(fileOpts) {
		log.Logger().Fatal("Shutting down due to configuration errors.")
	}

	log.SetLogger(log.ConfigureLogger().
		WithField(log.FieldSource, fileOpts.ServiceName))

	serve := server.NewServer(cliOpts, fileOpts)

	serve.RegisterEndpoints()
	log.Logger().Fatal(serve.Serve())
}

func checkErr(err error) {
	if err != nil {
		log.Logger().Fatal(err)
	}
}
