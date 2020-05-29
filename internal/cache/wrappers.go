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

func (m *Meta) Set(token string, data job.Metadata) {
	m.cache.Set(token, data, DefaultExpiration)
}

func (m *Meta) Get(token string) (out job.Metadata, ok bool) {
	tmp, ok := m.cache.Get(token)

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

func (u *Upload) SetDetails(token string, details job.Details) {
	u.cache.Set(token, details, DefaultExpiration)
}

func (u *Upload) SetStorable(token string, details job.StorableDetails) {
	u.cache.Set(token, details, DefaultExpiration)
}

func (u *Upload) GetDetails(token string) (out job.Details, ok bool) {
	tmp, ok := u.cache.Get(token)

	if !ok {
		return
	}

	out, ok = tmp.(job.Details)

	return
}

func (u *Upload) GetStorable(token string) (out job.StorableDetails, ok bool) {
	tmp, ok := u.cache.Get(token)

	if !ok {
		return
	}

	out, ok = tmp.(job.StorableDetails)

	return
}
