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

	router.Use(m.Tracing(), m.Translations(), m.Metrics(), m.Recovery())
	r0 := router.Group("")
	{
		r0.GET("/metrics", gin.WrapH(promhttp.Handler()))
	}
	r1 := router.Group("")
	{
		r1.GET("/v1/captcha", v1.Captcha) // 获取验证码
		//r1.POST("/v1/token", v1.Token)       // 获取token
		//r1.POST("/v1/register", v1.Register) // 注册
		r1.POST("/v1/login", v1.Login) // 登录
	}
	r2 := router.Group("").Use(m.JWTAuth(), m.Casbin())
	{
		r2.GET("/v1/getUserInfo", v1.UserInfo)          // 获取用户信息
		r2.POST("/v1/getUserInfoList", v1.UserInfoList) // 获取用户信息列表
		r2.POST("/v1/setUserInfo", v1.SetUserInfo)      // 更新用户信息
		r2.POST("/v1/getRoleMenus", v1.GetRoleMenus)    //
		//r2.POST("/addCasbin", v1.AddCasbin)                   // 添加权限
		//r2.POST("/addCasbinWithExcel", v1.AddCasbinWithExcel) // 批量添加权限（excel文件方式）
		//r2.POST("/getRoleMenus", v1.GetRoleMenus) // 获取角色路由
		//r2.POST("/getDict", v1.GetDict)
		//r2.POST("/upload", v1.GetDict)   //用户拼接
		//r2.POST("/download", v1.GetDict) //用户拼接
	}

	return router
}
