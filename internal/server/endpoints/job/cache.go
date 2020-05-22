package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/patrickmn/go-cache"
	"time"
)

func (e *endpoint) CreateDetails(meta *Metadata) *job.Details {
	details := job.Details{
		StorableDetails: job.StorableDetails{
			Started:  time.Now(),
			UserID:   meta.Owner,
			Token:    meta.Token,
			Status:   job.StatusReceiving,
			Projects: meta.Projects,
		},
	}
	e.upload.Set(meta.Token, details, cache.DefaultExpiration)
	return &details
}

func (e *endpoint) StoreDetails(details *job.Details) {
	e.upload.Set(details.Token, *details, cache.DefaultExpiration)
}
