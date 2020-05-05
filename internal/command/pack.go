package command

import (
	"errors"

	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

const errPackFail = "Failed to package archive for download: "

func (r *runner) packArchive(files []string) error {
	cmd := util.PrepCommand(r.log, "tar", "-cvzf", "dataset.tgz")
	cmd.Dir = r.details.WorkingDir
	cmd.Args = append(cmd.Args, files...)

	err := cmd.Run()
	if err != nil {
		err = errors.New(errPackFail + err.Error())
	}

	return err
}
