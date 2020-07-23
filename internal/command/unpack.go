package command

import (
	"errors"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/util/archive"
)

const (
	errUnzip = "failed to unzip uploaded archive: "
	errUntar = "failed to extract uploaded tar archive: "
)

// Unpack the uploaded archive file into the working directory.
//
// This method handles zip files, tar files, and raw biom files.
func (r *runner) unpack(d *job.Details) error {
	if archive.HasZipExtension(d.InputFile) {
		files, err := archive.UnZip(d.WorkingDir, d.InputFile, r.log)
		if err != nil {
			return errors.New(errUnzip + err.Error())
		}

		d.UnpackedFiles = files
	} else if archive.HasTarExtension(d.InputFile) {
		files, err := archive.UnTar(d.WorkingDir, d.InputFile, r.log)
		if err != nil {
			return errors.New(errUntar + err.Error())
		}

		d.UnpackedFiles = files
	} else {
		d.UnpackedFiles = []string{d.InputFile}
	}

	return nil
}
