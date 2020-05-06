package parse

import (
	// Std Lib
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/app/service"
	"github.com/VEuPathDB/util-exporter-server/internal/log"
	"os"

	// External
	"github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/app"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
)

const (
	argPort  = "port"
	helpPort = "Port the server should bind to"
	defPort  = uint16(80)

	argOpts  = "config"
	helpOpts = "Path to service configuration file.\n\n" +
		"Defaults to " + defOpts + " for containerized usage."
	defOpts  = "/app/config.yml"

	argWdir  = "workspace"
	helpWdir = "Path to workspace directory.  If this directory does not " +
		"already exist, it will be created."
	defWdir = "/workspace"

	argVer  = "version"
	helpVer = "Print server version"

	argMode  = "app-mode"
	helpMode = "Server app run mode.  Options are:\n\n" +
		modeServe + "        = run dataset processing HTTP server\n" +
		modeValidate + " = validate the server configuration file.\n" +
		modeGenerate + "   = generate an example configuration file\n\n" +
		"Default mode is \"serve\"."
)

const (
	modeServe    = "serve"
	modeValidate = service.ValidateAppName
	modeGenerate = service.GenerateAppName
)

func Cli(opts *config.Options) {
	var mode string

	cli.NewCommand().
		Flag(cli.LFlag(argPort, helpPort).
			Bind(&opts.Port, true).
			Default(defPort)).
		Flag(cli.LFlag(argOpts, helpOpts).
			Bind(&opts.ConfigPath, true).
			Default(defOpts)).
		Flag(cli.LFlag(argWdir, helpWdir).
			Bind(&opts.Workspace, true).
			Default(defWdir)).
		Flag(cli.LFlag(argVer, helpVer).
			OnHit(func(argo.Flag) {
				fmt.Println(opts.Version)
				os.Exit(app.StatusSuccess)
			})).
		Arg(cli.NewArg().
			Name(argMode).
			Description(helpMode).
			Default("serve").
			Bind(&mode)).
		MustParse()

	switch mode {
	case modeValidate:
		service.ValidateConfig(opts)
	case modeGenerate:
		service.GenerateConfig()
	case modeServe:
		// do nothing
	default:
		log.Logger().Fatal("invalid mode:", mode)
	}
}
