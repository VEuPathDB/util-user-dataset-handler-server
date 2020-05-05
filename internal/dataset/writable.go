package dataset

import "github.com/VEuPathDB/util-exporter-server/internal/wdk/site"

type Info struct {
	Owner        string         `json:"owner"`
	Projects     []site.WdkSite `json:"projects"`
	Type         Type           `json:"type"`
	Dependencies []Resource     `json:"dependencies"`
	Created      uint64         `json:"created"`
	Size         uint64         `json:"size"`
	DataFiles    []File         `json:"dataFiles"`
}

type Metadata struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Summary     string `json:"summary"`
}
