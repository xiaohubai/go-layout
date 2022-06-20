package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cors 跨域
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		origin := c.Request.Header.Get("Origin")       // 请求头部
		c.Header("Acess-Control-Allow-Origin", origin) // 允许所有请求
		// 运行跨域设置可以返回其他字段，可以自定义字段
		c.Header("Acess-Control-Allow-Headers", "Content-Type,Content-Length,Authorization,X-Request-ID,AccessToken")
		// 服务器支持的所有跨域请求的方法
		c.Header("Acess-Control-Allow-Methods", "POST, GET, OPTIONS,DELETE,PUT")
		//允许浏览器可以解析的头部
		c.Header("Acess-Control-Expose-Headers", "Content-Length,Acess-Control-Allow-Origin,Acess-Control-Allow-Headers,Content-Type")
		// 允许客户端传递校验信息比如 cookie
		c.Header("Acess-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		// 处理请求
		c.Next()
	}
}
