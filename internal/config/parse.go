package config

import (
	// Std Lib
	"fmt"
	"io/ioutil"
	"os"

	// External
	"github.com/Foxcapades/Argonaut/v0"
	"github.com/Foxcapades/Argonaut/v0/pkg/argo"
	"gopkg.in/yaml.v3"
)

const (
	argPort  = "port"
	helpPort = "Port the server should bind to"
	defPort  = uint16(80)

	argOpts  = "options"
	helpOpts = "Path to options.yml"
	defOpts  = "/app/options.yml"

	argWdir  = "workspace"
	helpWdir = "Path to workspace directory.  If this directory does not " +
		"already exist, it will be created."
	defWdir = "/workspace"

	argVer = "version"
	helpVer = "Print server version"
)

func ParseCli(opts *Options) {
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
				os.Exit(1)
			})).
		MustParse()
}

func ParseOptions(opts *Options) {
	raw, err := ioutil.ReadFile(opts.ConfigPath)
	if err != nil {
		panic(err)
	}

	tmp := new(Options)
	err = yaml.Unmarshal(raw, tmp)
	if err != nil {
		panic(err)
	}

	opts.Commands = tmp.Commands
	opts.ServiceName = tmp.ServiceName
}
