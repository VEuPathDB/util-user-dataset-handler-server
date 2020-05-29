package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	detailsLife  = 30 * time.Minute
	detailsClean = 15 * time.Minute
)

var progressCache = cache.New(detailsLife, detailsClean)

func GetDetails(jobID string) (job.Details, bool) {
	if tmp, ok := progressCache.Get(jobID); ok {
		return tmp.(job.Details), ok
	}

	return job.Details{}, false
}

func PutDetails(jobID string, details job.Details) {
	progressCache.SetDefault(jobID, details)
}

func DeleteDetails(jobID string) {
	progressCache.Delete(jobID)
}
