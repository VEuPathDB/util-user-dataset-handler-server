package dataset

import (
	"github.com/VEuPathDB/util-exporter-server/internal/wdk/site"
)

type BaseInfo struct {
	Type         Type           `json:"type"`
	Projects     []site.WdkSite `json:"projects"`
	Owner        uint           `json:"owner"`
	Dependencies []Resource     `json:"dependencies,omitempty"`
}

type Info struct {
	BaseInfo
	Size      uint   `json:"size"`
	Created   uint64 `json:"created"`
	DataFiles []File `json:"dataFiles"`
}

type Type struct {
	Version string `json:"version"`
	Name    string `json:"name"`
}

type File struct {
	File string `json:"file"`
	Name string `json:"name"`
	Size uint   `json:"size"`
}

type Resource struct {
	DisplayName string `json:"resourceDisplayName"`
	Version     string `json:"resourceVersion"`
	Identifier  string `json:"resourceIdentifier"`
}
