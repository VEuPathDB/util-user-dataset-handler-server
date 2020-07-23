package command

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/vulpine-io/split-pipe/v1/pkg/spipe"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/metrics"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

type response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Configure and run the given command.
func (r *runner) handleCommand(cmd config.Command) (err error) {
	r.log.Trace("command.runner.handleCommand")

	args, err := r.parseArgs(cmd.Args)
	if err != nil {
		return err
	}

	env := os.Environ()

	X := util.PrepCommand(r.log, cmd.Executable, args...)
	buffer := new(strings.Builder)
	X.Stderr = spipe.NewSplitWriter(X.Stderr, buffer)
	X.Dir = r.details.WorkingDir
	X.Env = env

	r.log.Debug("running command:", X)

	time, err := util.TimeCmd(X)
	metrics.RecordCommandTime(cmd.Executable, time)

	if err != nil {
		raw := strings.TrimSpace(buffer.String())
		obj := strings.IndexByte(raw, '{')

		if obj == -1 {
			return errors.New(strings.TrimSpace(raw))
		}

		tmp := response{}
		if err := json.Unmarshal([]byte(raw[obj:]), &tmp); err != nil {
			return errors.New(strings.TrimSpace(raw))
		}

		msg := tmp.Message
		if obj > 0 {
			msg += "\n" + raw[0:obj]
		}

		switch tmp.Error {
		case "user":
			return NewUserError(msg)
		default:
			return NewHandlerError(msg)
		}
	}

	return nil
}

// Parse the arguments configured in the command config and inject any template
// variables encountered.
func (r *runner) parseArgs(args []string) (out []string, err error) {
	r.log.Trace("command.runner.parseArgs")

	for _, fn := range config.InjectorList() {
		args, err = fn(&r.details, &r.meta, r.log).Inject(args)
		if err != nil {
			return
		}
	}

	return args, nil
}
