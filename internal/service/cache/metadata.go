package cache

import (
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/VEuPathDB/util-exporter-server/internal/job"
)

const (
	metaLife  = 5 * time.Minute
	metaClean = 5 * time.Minute
)

var metaCache = cache.New(metaLife, metaClean)

func HasMetadata(jobID string) bool {
	_, ok := metaCache.Get(jobID)
	return ok
}

func GetMetadata(jobID string) (job.Metadata, bool) {
	if tmp, ok := metaCache.Get(jobID); ok {
		return tmp.(job.Metadata), ok
	}

	return job.Metadata{}, false
}

func PutMetadata(jobID string, meta job.Metadata) {
	metaCache.SetDefault(jobID, meta)
}

func DeleteMetadata(jobID string) {
	metaCache.Delete(jobID)
}

func AllMetadata() map[string]cache.Item {
	return metaCache.Items()
}
