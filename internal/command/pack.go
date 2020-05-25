package command

import (
	"errors"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
)

const errPackFail = "Failed to package archive for response: "

var archiveFiles = []string{
	"datafiles",
	"dataset.json",
	"meta.json",
}

func (r *runner) packArchive() error {
	cmd := util.PrepCommand(r.log, "tar", "-cvzf", "dataset.tgz")
	cmd.Dir = r.details.WorkingDir
	cmd.Args = append(cmd.Args, archiveFiles...)

	if err := cmd.Run(); err != nil {
		return errors.New(errPackFail + err.Error())
	}

	return nil
}
