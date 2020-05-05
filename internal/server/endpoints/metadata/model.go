package metadata

import (
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/google/uuid"
)

type Metadata struct {
	dataset.BaseInfo

	Name        string `json:"name"`
	Summary     string `json:"summary"`
	Description string `json:"description"`
	Token       string `json:"token"`
}

func (M *Metadata) Validate() (out svc.ValidationResult) {
	if val := M.BaseInfo.Validate(); !val.Ok {
		out.Ok = false
		for k, v := range val.Result {
			out.Result[k] = v
		}
	}
	if len(M.Name) == 0 {
		out.AddError("name", "name is required")
	}
	if len(M.Token) == 0 {
		out.AddError("token", "token is required")
	} else if _, err := uuid.Parse(M.Token); err != nil {
		out.AddError("token", "token must be a valid UUID v4 string")
	}
	return
}
