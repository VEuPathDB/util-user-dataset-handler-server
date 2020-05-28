package cache

import (
	"github.com/VEuPathDB/util-exporter-server/internal/job"
	"github.com/patrickmn/go-cache"
)

const DefaultExpiration = cache.DefaultExpiration

func NewMeta(cache *cache.Cache) *Meta {
	return &Meta{cache}
}

type Meta struct {
	cache *cache.Cache
}

func (M *Meta) Set(token string, data job.Metadata) {
	M.cache.Set(token, data, DefaultExpiration)
}

func (M *Meta) Get(token string) (out job.Metadata, ok bool) {
	tmp, ok := M.cache.Get(token)

	if ok {
		out = tmp.(job.Metadata)
	}

	return
}

func NewUpload(cache *cache.Cache) *Upload {
	return &Upload{cache}
}

type Upload struct {
	cache *cache.Cache
}

func (U *Upload) SetDetails(token string, details job.Details) {
	U.cache.Set(token, details, DefaultExpiration)
}

func (U *Upload) SetStorable(token string, details job.StorableDetails) {
	U.cache.Set(token, details, DefaultExpiration)
}

func (U *Upload) GetDetails(token string) (out job.Details, ok bool) {
	tmp, ok := U.cache.Get(token)

	if !ok {
		return
	}

	out, ok = tmp.(job.Details)
	return
}

func (U *Upload) GetStorable(token string) (out job.StorableDetails, ok bool) {
	tmp, ok := U.cache.Get(token)

	if !ok {
		return
	}

	out, ok = tmp.(job.StorableDetails)
	return
}