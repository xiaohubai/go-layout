package middleware

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/tracing"
	"go.opentelemetry.io/otel/attribute"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx, span := tracing.NewSpan(c.Request.Context(), "api")
				c.Request = c.Request.WithContext(ctx)
				defer span.End()

				buf := make([]byte, 2048)
				buf = buf[:runtime.Stack(buf, false)]
				bufs := string(buf)

				span.SetAttributes(attribute.Key("painc").String(bufs))
				response.Fail(c, response.CommonFailed, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
