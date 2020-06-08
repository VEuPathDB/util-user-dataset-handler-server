package archive

import (
	"github.com/VEuPathDB/util-exporter-server/internal/metrics"
	"github.com/VEuPathDB/util-exporter-server/internal/util"
	"github.com/VEuPathDB/util-exporter-server/internal/util/fs"
	"github.com/VEuPathDB/util-exporter-server/internal/util/slice"
	"os"
	"os/exec"
	"sort"
)

func unpack(dir, file string, cmd *exec.Cmd) ([]string, error) {
	// Get the list of files that were in the directory before we unpack the
	// archive.  This will be used later to figure out what we actually unpacked.
	before, err := fs.ListFiles(dir)
	if err != nil {
		return nil, err
	}

	// Set the command working directory to the given dir.
	cmd.Dir = dir

	// Run & time the unpack command for metrics.
	time, err := util.TimeCmd(cmd)
	metrics.RecordCommandTime(cmd.Args[0], time)

	if err != nil {
		return nil, err
	}

	// Remove the archive file itself
	err = os.Remove(file)
	if err != nil {
		return nil, err
	}

	// List the files in the directory after the unpack.
	after, err := fs.ListFiles(dir)
	if err != nil {
		return nil, err
	}

	// Make a new list containing only entries that exist in the directory now and
	// were not in the directory last time we checked.
	out := make([]string, 0, len(after) - len(before))
	for i := range after {
		if !slice.ContainsString(before, after[i]) {
			out = append(out, after[i])
		}
	}

	// Sort the output list by name ascending
	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out, err
}