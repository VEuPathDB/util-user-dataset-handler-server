package cache

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	historyLife  = 72 * time.Hour
	historyClean = 2 * time.Hour
)

var historyCache = cache.New(historyLife, historyClean)

func GetHistoricalDetails(jobId string) (job.StorableDetails, bool) {
	if tmp, ok := historyCache.Get(jobId); ok {
		return tmp.(job.StorableDetails), ok
	}

	return job.StorableDetails{}, false
}

func PutHistoricalDetails(jobId string, details job.StorableDetails) {
	historyCache.SetDefault(jobId, details)
}
