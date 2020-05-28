package command

import (
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/config"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"os"
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

	X := util.PrepCommand(r.log, cmd.Executable, args...)
	X.Dir = r.details.WorkingDir
	X.Env = env

	err = X.Run()

	if err != nil {
		r.log.Debug(X)
		return fmt.Errorf(errComFail, cmd.Executable, err)
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
