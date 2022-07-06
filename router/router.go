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

	router.Use(m.Cors(), m.Tracing(), m.Translations(), m.Metrics(), m.Recovery())
	r0 := router.Group("")
	{
		r0.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}

	r1 := router.Group("v1")
	{
		r1.GET("/captcha", v1.Captcha)
		r1.POST("/token", v1.Token)
		r1.POST("/register", v1.Register)
		r1.POST("/login", v1.Login)
	}
	//认证+鉴权
	r2 := router.Group("v1").Use(m.JWTAuth(), m.Casbin())
	{
		r2.POST("/getUserInfo", v1.UserInfo)
		r2.POST("/addCasbin", v1.AddCasbin)
		r2.POST("/addCasbinWithExcel", v1.AddCasbinWithExcel)
		r2.POST("/getDict", v1.GetDict)
		r2.POST("/upload", v1.GetDict)   //用户拼接
		r2.POST("/download", v1.GetDict) //用户拼接
	}

	return router
}
