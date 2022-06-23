package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/model"
	"gorm.io/gorm"
)

func AddCasbin(c *gin.Context, t []model.CasbinRule) error {
	db := GetDB(c, "mysql").(*gorm.DB)
	return db.Create(&t).Error
}
