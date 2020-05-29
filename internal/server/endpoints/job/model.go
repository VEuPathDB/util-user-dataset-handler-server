package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
)

type Metadata struct {
	job.Metadata
}

func (m *Metadata) Validate() (out svc.ValidationResult) {
	out = m.BaseInfo.Validate()

	if len(m.Name) == 0 {
		out.AddError("name", "name is required")
	}

	if len(m.Projects) == 0 {
		out.AddError("projects", "at least one project is required")
	}

	return
}
