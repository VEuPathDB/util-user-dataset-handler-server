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

type errorOutput struct {
	// Error type.
	Error string `json:"error"`
	// Error message.
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

	// Execute the command wrapped in a timer to record execution metrics.
	time, err := util.TimeCmd(X)
	metrics.RecordCommandTime(cmd.Executable, time)

	// If the command exited with a non-zero code (err is not nil)
	if err != nil {
		// Get the raw stderr buffer and trim off any extra whitespace characters.
		raw := strings.TrimSpace(buffer.String())

		// Check to see if the output appears to contain a json object, record the
		// position of the opening bracket.
		obj := strings.IndexByte(raw, '{')

		// If an opening curly bracket was not found, exit here with just the raw
		// text as an error message.
		if obj == -1 {
			return errors.New(strings.TrimSpace(raw))
		}

		// Since the output appears to contain a json object, attempt to parse it.
		tmp := errorOutput{}
		// If there was an error while parsing, then it was not a valid json object,
		// or contained extra trailing characters.  Spit out the raw stderr text as
		// an error.
		if err := json.Unmarshal([]byte(raw[obj:]), &tmp); err != nil {
			return errors.New(strings.TrimSpace(raw))
		}

		// If we're here, then there was a json object at the end of the stderr
		// output.
		msg := tmp.Message
		// If there was extra text before the json object, prepend it to the error
		// message from the json object.
		if obj > 0 {
			msg += "\n" + raw[0:obj]
		}

		// If the error was a "user" error, return that error type, else consider
		// it a handler error and return _that_ type.
		switch tmp.Error {
		case "user":
			return NewUserError(msg)
		default:
			return NewHandlerError(msg)
		}
	}

	// If we're here, then the command executed successfully.
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
