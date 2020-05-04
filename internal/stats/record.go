package stats

import (
	"sync"
	"time"
)

type ServerStats interface {
	RecordTime(duration time.Duration)
	RecordSize(bytes uint)
	ToPublic() interface{}
}

func GetServerStatus() ServerStats {
	return &status
}

var status stats

func init() {
	status = stats{
		byStatus: make(map[int]uint),
	}
}

type stats struct {
	lock sync.RWMutex

	byStatus map[int]uint
	times    []time.Duration
	sizes    []uint
	longest  time.Duration
	largest  uint
}

func (s *stats) RecordTime(dur time.Duration) {
	s.lock.Lock()
	s.times = append(s.times, dur)
	if s.longest < dur {
		s.longest = dur
	}
	s.lock.Unlock()
}

func (s *stats) RecordSize(bytes uint) {
	s.lock.Lock()
	s.sizes = append(s.sizes, bytes)
	if bytes > s.largest {
		s.largest = bytes
	}
	s.lock.Unlock()
}

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

	status := make(map[int]uint, len(s.byStatus))

	for k, v := range s.byStatus {
		status[k] = v
	}

	return requests{
		Longest:     s.longest,
		Largest:     s.largest,
		AvgDuration: s.averageTime(),
		AvgSize:     s.averageSize(),
		ByStatus:    status,
	}
}
