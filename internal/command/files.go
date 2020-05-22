package command

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	errDirReadFail = "Failed to read workspace directory: "
	errStatReadFail = "Failed to read command status output: "
	errStatFileOld = "Command status output not generated from command %s"
	errStatFileFail = "Failed to stat output file %s: %s"
)

// Returns the list of files in the working directory
func (r *runner) getWorkspaceFiles() ([]string, error) {
	files, err := ioutil.ReadDir(r.details.WorkingDir)
	if err != nil {
		return nil, errors.New(errDirReadFail + err.Error())
	}
	out := make([]string, 0, len(files) - 3)

	for _, file := range files {
		if file.Name() == "." || file.Name() == ".." || file.Name() == r.details.InTarName {
			continue
		}
		out = append(out, file.Name())
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out, nil
}

// getStatusFile reads the status file created by the
// previous command output.
func (r *runner) getStatusFile() (out job.StatusFile, err error) {
	file := path.Join(r.details.WorkingDir, "command-status.json")

	stat, err := os.Stat(file)
	if err != nil {
		return out, errors.New(errStatReadFail + err.Error())
	}

	// Verify that the file we are reading is newer than the
	// last file we read.  This should prevent the case where
	// a secondary command does not produce an output file.
	if !stat.ModTime().After(r.lastStatus) {
		return out, fmt.Errorf(errStatFileOld + r.lastCommand.Command)
	}

	bytes, err := ioutil.ReadFile(file)
	if err == nil {
		err = json.Unmarshal(bytes, &out)
	}

	if err != nil {
		err = errors.New(errStatReadFail + err.Error())
		return
	}

	return
}

func (r *runner) getFileSize(files []string) ([]uint64, error) {
	total := make([]uint64, len(files))
	for _, file := range files {
		stat, err := os.Stat(path.Join(r.details.WorkingDir, file))
		if err != nil {
			return nil, fmt.Errorf(errStatFileFail, file, err.Error())
		}
		total = append(total, uint64(stat.Size()))
	}
	return total, nil
}