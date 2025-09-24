package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"net/http"
)

var (
	// Counter: how many HTTP requests
	HttpRequestsTotal = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "Total number of HTTP requests",
		},
		[]string{"method", "path", "status"},
	)

	// Histogram: request duration
	HttpRequestDuration = promauto.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "http_request_duration_seconds",
			Help:    "Duration of HTTP requests in seconds",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"method", "path"},
	)

	StudentCacheHits = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "student_cache_hits_total",
			Help: "Total number of student cache hits",
		},
		[]string{"operation"},
	)

	// Cache Misses
	StudentCacheMisses = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "student_cache_misses_total",
			Help: "Total number of student cache misses",
		},
		[]string{"operation"},
	)
)

// Expose /metrics endpoint
func RegisterMetricsEndpoint() {
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(":2112", nil) // Prometheus scrapes here
}
