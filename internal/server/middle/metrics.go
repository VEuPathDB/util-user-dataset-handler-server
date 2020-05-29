package middle

import (
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"net/http"
	"strconv"
	"time"

	"github.com/Foxcapades/go-midl/v2/pkg/midl"
	"github.com/VEuPathDB/util-exporter-server/internal/service/logger"
)

var (
	promTotalRequests = promauto.NewCounterVec(prometheus.CounterOpts{
		Namespace:   "http",
		Name:        "total_requests",
		Help:        "Total HTTP request count.",
	}, []string{"path", "method", "status"})

	promRequestTime = promauto.NewHistogramVec(prometheus.HistogramOpts{
		Namespace:   "http",
		Name:        "request_duration",
		Help:        "Request times in milliseconds",
		Buckets:     []float64{0.05, 0.1, 0.5, 1, 5, 10},
	}, []string{"path", "method"})
)

func MetricAgg(next http.Handler) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		start := time.Now()
		wrap := resWrap{real: res}
		next.ServeHTTP(&wrap, req)
		q, _ := midl.NewRequest(req)
		recordTime(start, q, wrap.code)
	}
}

func recordTime(start time.Time, req midl.Request, code int) {
	dur := time.Since(start)
	met := req.RawRequest().Method
	url, _ := mux.CurrentRoute(req.RawRequest()).GetPathTemplate()

	logger.ByRequest(req).
		WithField("duration", dur.String()).
		WithField("status", code).
		Info("Request completed")

	promTotalRequests.WithLabelValues(url, met, strconv.Itoa(code)).Inc()
	promRequestTime.WithLabelValues(url, met).Observe(
		float64(dur) / float64(time.Millisecond))
}

type resWrap struct {
	code int
	real http.ResponseWriter
}

func (r *resWrap) Header() http.Header {
	return r.real.Header()
}

func (r *resWrap) Write(bytes []byte) (int, error) {
	return r.real.Write(bytes)
}

func (r *resWrap) WriteHeader(statusCode int) {
	r.code = statusCode
	r.real.WriteHeader(statusCode)
}
