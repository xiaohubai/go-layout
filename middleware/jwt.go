package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/jwt"
)

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.Split(c.Request.Header.Get("Authorization"), " ")
		if len(auth) != 2 {
			response.Fail(c, response.TokenFailed, nil)
			c.Abort()
			return
		}
		token := auth[1]
		claims, err := jwt.ParseToken(token)
		if err != nil {
			if err == consts.TokenExpired {
				response.Fail(c, response.TokenExpired, nil)
				c.Abort()
				return
			}
			response.Fail(c, response.TokenFailed, nil)
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
