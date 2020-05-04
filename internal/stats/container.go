package stats

import "time"

type container struct {
	Requests requests `json:"requests"`
}

type requests struct {
	Longest     *requestDetails `json:"longest,omitempty"`
	Largest     *requestDetails `json:"largest,omitempty"`
	AvgDuration time.Duration   `json:"avgDuration"`
	AvgSize     uint            `json:"avgSize"`
	ByStatus    map[int]uint    `json:"byStatus"`
}

type requestDetails struct {
	UserID   uint          `json:"userId"`
	Duration time.Duration `json:"duration"`
	Started  time.Time     `json:"started"`
	Size     uint          `json:"size"`
}
