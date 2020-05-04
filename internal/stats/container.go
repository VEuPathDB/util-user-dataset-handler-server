package stats

import "time"

type container struct {
	Requests requests `json:"requests"`
}

type requests struct {
	Longest     time.Duration `json:"longest"`
	Largest     uint          `json:"largest"`
	AvgDuration time.Duration `json:"avgDuration"`
	AvgSize     uint          `json:"avgSize"`
	ByStatus    map[int]uint  `json:"byStatus"`
}
