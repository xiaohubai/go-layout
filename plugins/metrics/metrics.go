package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	prometheus.MustRegister(PathCounter, KeyCounter)
}

var PathCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "go_layout_request_url_total",
	},
	[]string{"method", "url"},
)

var KeyCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "go_layout_key_total",
	},
	[]string{"key", "datetime"},
)

func Counter(key string) prometheus.Counter {
	return KeyCounter.With(prometheus.Labels{
		"datetime": time.Now().Format("2006-01-02 15:04:05"),
		"key":      key,
	})
}
