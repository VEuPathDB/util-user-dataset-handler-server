package dataset

import (
	"fmt"
	"github.com/VEuPathDB/util-exporter-server/internal/server/svc"
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
	"strconv"
)

type BaseInfo struct {
	Projects     []site.WdkSite `json:"projects"`
	Owner        uint           `json:"owner"`
	Dependencies []Resource     `json:"dependencies,omitempty"`
}

func (B *BaseInfo) Validate() (out svc.ValidationResult) {

	for i := range B.Projects {
		if !B.Projects[i].IsValid() {
			out.AddError("projects", fmt.Sprintf("unrecognized project id '%s'",
				B.Projects[i]))
		}
	}

	if B.Owner == 0 {
		out.AddError("owner", "owner is required")
	}

	for i := range B.Dependencies {
		if val := B.Dependencies[i].Validate(); val.Failed {
			base := "dependencies[" + strconv.Itoa(i) + "]."
			for k, v := range val.Result {
				out.Result[base + k] = v
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

func (R *Resource) Validate() (out svc.ValidationResult) {
	if len(R.DisplayName) == 0 {
		out.AddError("resourceDisplayName", "resource display name is required")
	}
	if len(R.Version) == 0 {
		out.AddError("resourceVersion", "resource version is required")
	}
	if len(R.Identifier) == 0 {
		out.AddError("resourceIdentifier", "resource identifier is required")
	}
	return
}
