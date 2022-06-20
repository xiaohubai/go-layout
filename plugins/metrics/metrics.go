package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

func Init() {
	prometheus.MustRegister(PathCounter, KeyCounter, CounterM)
}

var PathCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "each_url_total",
	},
	[]string{"method", "path"},
)

var KeyCounter = prometheus.NewCounterVec(
	prometheus.CounterOpts{
		Name: "key_total",
	},
	[]string{"key"},
)

var CounterM = prometheus.NewHistogramVec(
	prometheus.HistogramOpts{
		Name: "minute_key",
		Buckets: []float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13,
			14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28,
			29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43,
			44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59},
	},
	[]string{"date", "key"},
)

func CounterInc(key string) {
	KeyCounter.With(prometheus.Labels{
		"key": key,
	}).Inc()
}

func CounterIncM(key string) {
	CounterM.With(prometheus.Labels{
		"date": time.Now().Format("2006-01-02 15"),
		"key":  key,
	}).Observe(float64(time.Now().Minute()))
}
