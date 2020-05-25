package command

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/VEuPathDB/util-exporter-server/internal/config"
)

const (
	errComFail = `Command "%s" failed: %s`
)

// Configure and run the given command.
func (r *runner) handleCommand(cmd *config.Command) (err error) {
	args, err := r.parseArgs(cmd.Args)
	if err != nil {
		return err
	}

	env := os.Environ()

	X := exec.Command(cmd.Executable)
	X.Args = args
	X.Env  = env

	err = X.Run()

	if err != nil {
		return fmt.Errorf(errComFail, cmd.Executable, err)
	}

	return nil
}

// Parse the arguments configured in the command config and inject any template
// variables encountered.
func (r *runner) parseArgs(args []string) (out []string, err error) {
	for _, fn := range config.InjectorList() {
		args, err = fn(&r.details).Inject(args)
		if err != nil {
			return
		}
	}

	return args, nil
}
