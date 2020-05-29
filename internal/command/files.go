package command

import (
	"os"
	"sort"
)

// Returns the list of files in the working directory.
func (r *runner) getWorkspaceFiles() ([]string, error) {
	files, err := r.wkspc.Files(func(info os.FileInfo) bool {
		return info.Name() != r.details.InTarName
	})

	if err != nil {
		return nil, err
	}

	out := make([]string, len(files))

	for i, file := range files {
		out[i] = file.Name()
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out, nil
}
