package stats

import (
	"github.com/VEuPathDB/util-exporter-server/internal/process"
	"sync"
	"time"
)

type stats struct {
	lock sync.RWMutex

	byStatus map[int]uint
	times    []time.Duration
	sizes    []uint
	longest  *requestDetails
	largest  *requestDetails
}

func (s *stats) RecordTime(dur time.Duration, details *process.StorableDetails) {
	s.lock.Lock()

	s.lock.Unlock()
}

func (s *stats) recDetails()

func (s *stats) averageSize() uint {
	total := uint64(0)
	for _, v := range s.sizes {
		total += uint64(v)
	}
	return uint(total / uint64(len(s.sizes)))
}

func (s *stats) averageTime() time.Duration {
	total := time.Duration(0)
	for _, v := range s.times {
		total += v
	}
	return total / time.Duration(len(s.times))
}

func (s *stats) ToPublic() interface{} {
	s.lock.RLock()
	defer s.lock.RUnlock()

	tmpLong := *s.longest
	tmpLarge := *s.largest
	status := make(map[int]uint, len(s.byStatus))

	for k, v := range s.byStatus {
		status[k] = v
	}

	return container{
		Requests: requests{
			Longest:     &tmpLong,
			Largest:     &tmpLarge,
			AvgDuration: s.averageTime(),
			AvgSize:     s.averageSize(),
			ByStatus:    status,
		},
	}
}
