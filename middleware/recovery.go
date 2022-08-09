package middleware

import (
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/xiaohubai/go-layout/model/response"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
				c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
				defer span.Finish()

				buf := make([]byte, 2048)
				buf = buf[:runtime.Stack(buf, false)]
				bufs := string(buf)

				span.LogFields(log.Object("Recovery()", err), log.Object("error", bufs))
				response.Fail(c, response.CommonFailed, nil)
				c.Abort()
			}
		}()
		c.Next()
	}
}
