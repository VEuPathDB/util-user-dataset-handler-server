package command

import (
	"encoding/json"
	"errors"
	"os"
	"strings"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

type response struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

// Configure and run the given command.
func (r *runner) handleCommand(cmd *config.Command) (err error) {
	args, err := r.parseArgs(cmd.Args)
	if err != nil {
		return err
	}

	env := os.Environ()

	X := util.PrepCommand(r.log, cmd.Executable, args...)
	buffer := util.NewBufferPipe(X.Stderr)
	X.Stderr = buffer
	X.Dir = r.details.WorkingDir
	X.Env = env

	time, err := util.TimeCmd(X)
	promCommandTime.WithLabelValues(cmd.Executable).Observe(time)

	if err != nil {
		r.log.Debug(X)

		raw := strings.TrimSpace(buffer.Buffer.String())
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
	for _, fn := range config.InjectorList() {
		args, err = fn(&r.details, &r.meta).Inject(args)
		if err != nil {
			return
		}
	}

	return args, nil
}
