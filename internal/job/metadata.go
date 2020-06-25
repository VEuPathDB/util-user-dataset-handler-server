package job

import "github.com/VEuPathDB/util-exporter-server/internal/dataset"

type Metadata struct {
	dataset.BaseInfo

	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Token       string `json:"jobId"`
	Origin      string `json:"origin"`
}
