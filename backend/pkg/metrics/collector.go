package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

// Collector 指标收集器
type Collector struct {
	RequestCounter   *prometheus.CounterVec
	RequestDuration  *prometheus.HistogramVec
	ResponseSize     *prometheus.HistogramVec
	ActiveConnGauge  prometheus.Gauge
	ErrorCounter     *prometheus.CounterVec
	CacheHitCounter  *prometheus.CounterVec
	CacheMissCounter *prometheus.CounterVec
}

// NewCollector 创建新的指标收集器
func NewCollector() *Collector {
	return &Collector{
		RequestCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "http_requests_total",
				Help: "Total number of HTTP requests",
			},
			[]string{"method", "path", "status"},
		),
		RequestDuration: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_request_duration_seconds",
				Help:    "HTTP request duration in seconds",
				Buckets: prometheus.DefBuckets,
			},
			[]string{"method", "path"},
		),
		ResponseSize: promauto.NewHistogramVec(
			prometheus.HistogramOpts{
				Name:    "http_response_size_bytes",
				Help:    "HTTP response size in bytes",
				Buckets: prometheus.ExponentialBuckets(100, 10, 8),
			},
			[]string{"method", "path"},
		),
		ActiveConnGauge: promauto.NewGauge(
			prometheus.GaugeOpts{
				Name: "http_active_connections",
				Help: "Number of active HTTP connections",
			},
		),
		ErrorCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "application_errors_total",
				Help: "Total number of application errors",
			},
			[]string{"type", "code"},
		),
		CacheHitCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "cache_hits_total",
				Help: "Total number of cache hits",
			},
			[]string{"cache"},
		),
		CacheMissCounter: promauto.NewCounterVec(
			prometheus.CounterOpts{
				Name: "cache_misses_total",
				Help: "Total number of cache misses",
			},
			[]string{"cache"},
		),
	}
}

// RecordRequest 记录HTTP请求
func (c *Collector) RecordRequest(method, path string, status int) {
	c.RequestCounter.WithLabelValues(method, path, string(status)).Inc()
}

// RecordRequestDuration 记录请求持续时间
func (c *Collector) RecordRequestDuration(method, path string, duration float64) {
	c.RequestDuration.WithLabelValues(method, path).Observe(duration)
}

// RecordResponseSize 记录响应大小
func (c *Collector) RecordResponseSize(method, path string, size float64) {
	c.ResponseSize.WithLabelValues(method, path).Observe(size)
}

// RecordActiveConnection 记录活动连接数
func (c *Collector) RecordActiveConnection(delta float64) {
	c.ActiveConnGauge.Add(delta)
}

// RecordError 记录错误
func (c *Collector) RecordError(errorType, code string) {
	c.ErrorCounter.WithLabelValues(errorType, code).Inc()
}

// RecordCacheHit 记录缓存命中
func (c *Collector) RecordCacheHit(cache string) {
	c.CacheHitCounter.WithLabelValues(cache).Inc()
}

// RecordCacheMiss 记录缓存未命中
func (c *Collector) RecordCacheMiss(cache string) {
	c.CacheMissCounter.WithLabelValues(cache).Inc()
}
