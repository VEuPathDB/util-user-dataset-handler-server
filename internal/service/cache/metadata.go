package cache

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/patrickmn/go-cache"
	"time"
)

const (
	metaLife  = 5 * time.Minute
	metaClean = 5 * time.Minute
)

var metaCache = cache.New(metaLife, metaClean)

func HasMetadata(jobId string) bool {
	_, ok := metaCache.Get(jobId)
	return ok
}

func GetMetadata(jobId string) (job.Metadata, bool) {
	if tmp, ok := metaCache.Get(jobId); ok {
		return tmp.(job.Metadata), ok
	}

	return job.Metadata{}, false
}

func PutMetadata(jobId string, meta job.Metadata) {
	metaCache.SetDefault(jobId, meta)
}

func DeleteMetadata(jobId string) {
	metaCache.Delete(jobId)
}