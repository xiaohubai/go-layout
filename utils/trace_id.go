package utils

import (
	"github.com/gin-gonic/gin"
)

func TraceId(c *gin.Context) string {
	if id, ok := c.Get("X-Trace-ID"); ok {
		return id.(string)
	}
	return ""
}
