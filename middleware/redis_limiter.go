package middleware

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis_rate/v9"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/utils"
	"go.uber.org/zap"
)

func RedisLimiter() gin.HandlerFunc {
	return func(c *gin.Context) {
		rdb := global.Redis
		limiter := redis_rate.NewLimiter(rdb)
		uri := c.Request.RequestURI
		index := strings.Index(uri, "?")
		var key string
		if index == -1 {
			key = uri
		} else {
			key = uri[:index]
		}
		res, err := limiter.Allow(c, key, redis_rate.PerMinute(global.Cfg.System.Rate))
		if err != nil {
			global.Log.Error(utils.TraceId(c), zap.Any("key", "mid"), zap.Any("msg", fmt.Sprintf("%s:%s", "分布式限流", err)))
		} else {
			c.Header("RateLimit-Remaining", strconv.Itoa(res.Remaining))
			if res.Allowed == 0 {
				seconds := int(res.RetryAfter / time.Second)
				c.Header("RateLimit-RetryAfter", strconv.Itoa(seconds))
				response.Fail(c, response.ErrRateLimited, nil)
				c.Abort()
			}
		}
		c.Next()
	}
}
