package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	promCommandTime = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: "cmd",
			Name:      "execution_time",
			Help:      "Execution time for external command calls in milliseconds.",
			// 50ms, 100ms, 1s, 10s, 1m, 5m, 10m
			Buckets: []float64{50, 100, 1000, 10_000, 60_000, 300_000, 600_000},
		},
		[]string{"command"},
	)
)

func RecordCommandTime(cmd string, time float64) {
	promCommandTime.WithLabelValues(cmd).Observe(time)
}
