package cache

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	detailsLife  = 30 * time.Minute
	detailsClean = 15 * time.Minute
)

var progressCache = cache.New(detailsLife, detailsClean)

func GetDetails(jobId string) (job.Details, bool) {
	if tmp, ok := progressCache.Get(jobId); ok {
		return tmp.(job.Details), ok
	}

	return job.Details{}, false
}

func PutDetails(jobId string, details job.Details) {
	progressCache.SetDefault(jobId, details)
}

func DeleteDetails(jobId string) {
	progressCache.Delete(jobId)
}
