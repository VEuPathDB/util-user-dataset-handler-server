package dataset

import (
	"fmt"

	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
)

type BaseInfo struct {
	Projects     []site.WdkSite `json:"projects"`
	Owner        uint           `json:"owner"`
	Dependencies []Resource     `json:"dependencies,omitempty"`
}

func (b *BaseInfo) Validate() (out svc.ValidationResult) {
	for i := range b.Projects {
		if !b.Projects[i].IsValid() {
			out.AddError("projects", fmt.Sprintf("unrecognized project id '%s'",
				b.Projects[i]))
		}
	}

	if b.Owner == 0 {
		out.AddError("owner", "owner is required")
	}

	for i := range b.Dependencies {
		if val := b.Dependencies[i].Validate(); val.Failed {
			base := fmt.Sprintf("dependencies[%d].", i)

			for k := range val.Result {
				out.Result[base+k] = val.Result[k]
			}
		}
	}

	return
}

type File struct {
	File string `json:"file"`
	Name string `json:"name"`
	Size uint64 `json:"size"`
}

type Resource struct {
	DisplayName string `json:"resourceDisplayName"`
	Version     string `json:"resourceVersion"`
	Identifier  string `json:"resourceIdentifier"`
}

func (r *Resource) Validate() (out svc.ValidationResult) {
	if len(r.DisplayName) == 0 {
		out.AddError("resourceDisplayName", "resource display name is required")
	}

	if len(r.Version) == 0 {
		out.AddError("resourceVersion", "resource version is required")
	}

	if len(r.Identifier) == 0 {
		out.AddError("resourceIdentifier", "resource identifier is required")
	}

	return
}
