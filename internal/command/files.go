package command

import (
	"sort"
)

// Returns the list of files in the working directory.
func (r *runner) getWorkspaceFiles() ([]string, error) {
	files, err := r.wkspc.Files(func(string) bool { return true })

	if err != nil {
		return nil, err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] < files[j]
	})

	return files, nil
}
