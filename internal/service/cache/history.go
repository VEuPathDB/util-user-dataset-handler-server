package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	historyLife  = 72 * time.Hour
	historyClean = 2 * time.Hour
)

var historyCache = cache.New(historyLife, historyClean)

func GetHistoricalDetails(jobID string) (job.StorableDetails, bool) {
	if tmp, ok := historyCache.Get(jobID); ok {
		return tmp.(job.StorableDetails), ok
	}

	return job.StorableDetails{}, false
}

func PutHistoricalDetails(jobID string, details job.StorableDetails) {
	historyCache.SetDefault(jobID, details)
}

func AllHistoricalDetails() map[string]cache.Item {
	return historyCache.Items()
}
