package router

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	v1 "github.com/xiaohubai/go-layout/api/v1"
	m "github.com/xiaohubai/go-layout/middleware"
)

// Routers 初始化路由
func Routers() *gin.Engine {
	var router = gin.Default()

	router.Use(m.Cors(), m.Translations(), m.Metrics())

	r0 := router.Group("")
	{
		r0.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	r1 := router.Group("v1").Use(m.RedisLimiter())
	{
		r1.GET("/captcha", v1.Captcha)
		r1.POST("/token", v1.Token)
		r1.POST("/register", v1.Register)
		r1.POST("/login", v1.Login)
	}

	r2 := router.Group("v1").Use(m.JWTAuth()).Use(m.Casbin(), m.Tracing())
	{
		r2.POST("/getUserInfo", v1.UserInfo)
	}

	return router
}
