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
	before, err := fs.ListFiles(dir)
	if err != nil {
		return nil, err
	}

	cmd.Dir = dir

	time, err := util.TimeCmd(cmd)
	metrics.RecordCommandTime(cmd.Args[0], time)

	if err != nil {
		return nil, err
	}

	err = os.Remove(file)
	if err != nil {
		return nil, err
	}

	after, err := fs.ListFiles(dir)
	if err != nil {
		return nil, err
	}

	out := make([]string, 0, len(after) - len(before))
	for i := range after {
		if !slice.ContainsString(before, after[i]) {
			out = append(out, after[i])
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i] < out[j]
	})

	return out, err
}