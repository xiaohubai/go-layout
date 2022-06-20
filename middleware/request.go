package middleware

import (
	"bytes"
	"context"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/utils"

	"go.uber.org/zap"
)

func Request() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.RequestURI
		ip := c.ClientIP()
		mothod := c.Request.Method
		uid := utils.TraceId(c)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), "trace_id", uid))
		reqBody, _ := c.GetRawData()
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(reqBody))
		global.Log.Info(uid, zap.Any("key", "req"), zap.Any("ip", ip), zap.Any("path", path), zap.String("method", mothod), zap.Any("query", string(reqBody)))

		c.Next()

		respBpdy, _ := c.Get("resp_body")
		elapsed := time.Since(start)
		global.Log.Info(uid, zap.Any("key", "resp"), zap.Duration("elapsed", elapsed), zap.Any("body", respBpdy))
	}
}
