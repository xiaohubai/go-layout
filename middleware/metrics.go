package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/xiaohubai/go-layout/plugins/metrics"
)

func Metrics() gin.HandlerFunc {
	return func(c *gin.Context) {
		metrics.PathCounter.With(prometheus.Labels{
			"method": c.Request.Method,
			"url":    c.Request.RequestURI,
		}).Inc()

		c.Next()
	}
}
