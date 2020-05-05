package command

import (
	"errors"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"strings"

	"github.com/VEuPathDB/util-exporter-server/internal/process"
)

const (
	errUnzip   = "Failed to unzip uploaded archive: "
	errUntar   = "Failed to extract uploaded tar archive: "
)

func (r *runner) unpack(d *process.Details) error {
	if strings.HasSuffix(d.InTarName, ".zip") {
		if err := r.unzip(d); err != nil {
			return err
		}
	} else {
		if err := r.untar(d); err != nil {
			return err
		}
	}

	files, err := r.getWorkspaceFiles()
	if err != nil {
		return err
	}
	r.details.InputFiles = files
	return nil
}

func (r *runner) untar(d *process.Details) error {
	cmd := util.PrepCommand(r.log, "tar", "-xvf", d.InTarName)
	cmd.Dir = d.WorkingDir

	err := cmd.Run()
	if err != nil {
		err = errors.New(errUntar + err.Error())
	}
	return err
}

func (r *runner) unzip(d *process.Details) error {
	cmd := util.PrepCommand(r.log, "unzip", d.InTarName)
	cmd.Dir = d.WorkingDir

	err := cmd.Run()
	if err != nil {
		err = errors.New(errUnzip + err.Error())
	}
	return err
}