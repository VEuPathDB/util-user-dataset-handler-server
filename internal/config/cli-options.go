package config

import (
	"fmt"
	"os"
	"strconv"

	cli "github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"

	"github.com/VEuPathDB/util-exporter-server/internal/app"
	"github.com/VEuPathDB/util-exporter-server/pkg/meta"
)

const (
	argPort         = "port"
	helpPort        = "Port the server should bind to"
	defPort  uint16 = 80

	argOpts  = "config"
	helpOpts = "Path to service configuration file.\n\n" +
		"Defaults to " + defOpts + " for containerized usage."
	defOpts = "/app/config.yml"

	argWdir  = "workspace"
	helpWdir = "Path to workspace directory.  If this directory does not " +
		"already exist, it will be created."
	defWdir = "/workspace"

	argVer  = "version"
	helpVer = "Print server version"
)

// CLIOptions contains the values of the passed command line args (or the
// default values for those args).
type CLIOptions interface {
	// Port returns the configured port number the HTTP server should bind to.
	// Defaults to "80".
	Port() uint16

	// ConfigPath returns the configured path to the service config file.
	// Defaults to "/app/config.yml".
	ConfigPath() string

	// WorkspacePath returns the configured workspace root directory.
	// Defaults to "/workspace"
	WorkspacePath() string

	// GetUsablePort returns the configured server port in the format expected by
	// the Golang HTTP server package.
	GetUsablePort() string
}

func ParseCLIOptions() (CLIOptions, error) {
	out := new(cliOpts)

	_, err := cli.NewCommand().
		Flag(cli.LFlag(argPort, helpPort).
			Bind(&out.port, true).
			Default(defPort)).
		Flag(cli.LFlag(argOpts, helpOpts).
			Bind(&out.configPath, true).
			Default(defOpts)).
		Flag(cli.LFlag(argWdir, helpWdir).
			Bind(&out.workspacePath, true).
			Default(defWdir)).
		Flag(cli.LFlag(argVer, helpVer).
			OnHit(func(argo.Flag) {
				fmt.Println(meta.GetBuildMeta().String())
				os.Exit(app.StatusSuccess)
			})).
		Parse()

	return out, err
}

type cliOpts struct {
	port          uint16
	configPath    string
	workspacePath string
}

func (c *cliOpts) Port() uint16 {
	return c.port
}

func (c *cliOpts) ConfigPath() string {
	return c.configPath
}

func (c *cliOpts) WorkspacePath() string {
	return c.workspacePath
}

func (c *cliOpts) GetUsablePort() string {
	return ":" + strconv.Itoa(int(c.port))
}
