package options

import (
	// Std Lib
	"net/http"

	// External
	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/gorilla/mux"

	// Internal
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/server/middle"
)

const path = "/config"

// Register appends the configuration printout endpoint to the given router.
func Register(r *mux.Router, cli config.CLIOptions, file config.FileOptions) {
	r.Path(path).
		Methods(http.MethodGet).
		Handler(middle.MetricAgg(middle.RequestCtxProvider(
			midl.JSONAdapter(&configEndpoint{cli: cli, file: file}))))
}

type configEndpoint struct {
	cli  config.CLIOptions
	file config.FileOptions
}

func (c *configEndpoint) Handle(midl.Request) midl.Response {
	return midl.MakeResponse(http.StatusOK, output{
		ServiceName: c.file.ServiceName(),
		Command:     c.file.Commands(),
		FileTypes:   c.file.FileTypes(),
		ConfigPath:  c.cli.ConfigPath(),
		Workspace:   c.cli.WorkspacePath(),
	})
}

type output struct {
	ServiceName string         `json:"serviceName"`
	Command     config.Command `json:"command"`
	FileTypes   []string       `json:"fileTypes"`
	ConfigPath  string         `json:"configPath"`
	Workspace   string         `json:"workspace"`
}
