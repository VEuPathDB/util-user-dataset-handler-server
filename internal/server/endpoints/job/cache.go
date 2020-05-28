package job

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"time"
)

func (e *endpoint) CreateDetails(meta *job.Metadata) *job.Details {
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
	e.upload.SetDetails(meta.Token, details)
	return &details
}

func (e *endpoint) StoreDetails(details *job.Details) {
	e.upload.SetDetails(details.Token, *details)
}
