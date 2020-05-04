package xio

import (
	"github.com/VEuPathDB/util-exporter-server/internal/dataset"
)

type Metadata struct {
	dataset.BaseInfo

	Token string `json:"token"`
}
