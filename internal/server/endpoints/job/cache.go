package job

import (
	"time"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/VEuPathDB/util-exporter-server/internal/service/cache"
)

func (e *uploadEndpoint) createDetails(meta *job.Metadata) *job.Details {
	now := time.Now()
	details := job.Details{
		StorableDetails: job.StorableDetails{
			Started:  &now,
			UserID:   meta.Owner,
			Token:    meta.Token,
			Status:   job.StatusReceiving,
			Projects: meta.Projects,
		},
	}
	cache.PutDetails(meta.Token, details)

	return &details
}

func (e *uploadEndpoint) storeDetails(details *job.Details) {
	cache.PutDetails(details.Token, *details)
}
