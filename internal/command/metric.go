package command

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	promCommandTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace: "cmd",
		Name:      "execution_time",
		Help:      "Execution time for external command calls in milliseconds.",
		Buckets:   []float64{1, 10, 50, 100, 500, 1000, 5000},
	}, []string{"command"})
)
