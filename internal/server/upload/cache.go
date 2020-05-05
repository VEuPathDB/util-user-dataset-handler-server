package upload

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"github.com/VEuPathDB/util-exporter-server/internal/server/metadata"
	"github.com/patrickmn/go-cache"
	"time"
)

func (e *endpoint) CreateDetails(meta *metadata.Metadata) *process.Details {
	details := process.Details{
		StorableDetails: process.StorableDetails{
			Started:  time.Now(),
			UserID:   meta.Owner,
			Token:    meta.Token,
			Status:   process.StatusReceiving,
			Projects: meta.Projects,
		},
	}
	e.upload.Set(meta.Token, details, cache.DefaultExpiration)
	return &details
}

func (e *endpoint) StoreDetails(details *process.Details) {
	e.upload.Set(details.Token, *details, cache.DefaultExpiration)
}
