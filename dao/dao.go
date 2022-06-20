package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/configs/global"
)

func GetDB(c *gin.Context, name string) interface{} {
	switch name {
	case "mysql":
		return global.Db.WithContext(c.Request.Context())
	default:
		return global.Db.WithContext(c.Request.Context())
	}
}
