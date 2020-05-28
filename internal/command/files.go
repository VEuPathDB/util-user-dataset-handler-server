package command

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"sort"
)

const (
	errDirReadFail = "Failed to read workspace directory: "
	errStatFileFail = "Failed to stat output file %s: %s"
)

// Returns the list of files in the working directory
func (r *runner) getWorkspaceFiles() ([]string, error) {
	files, err := ioutil.ReadDir(r.details.WorkingDir)

	if err != nil {
		return nil, errors.New(errDirReadFail + err.Error())
	}

	out := make([]string, 0, len(files))

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