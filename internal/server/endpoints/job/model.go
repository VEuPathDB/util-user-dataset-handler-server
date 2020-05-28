package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

type Metadata struct {
	job.Metadata
}

func (M *Metadata) Validate() (out svc.ValidationResult) {
	out = M.BaseInfo.Validate()

	if len(M.Name) == 0 {
		out.AddError("name", "name is required")
	}

	if len(M.Projects) == 0 {
		out.AddError("projects", "at least one project is required")
	}

	return
}
