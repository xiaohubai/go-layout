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

func CasbinList(c *gin.Context, t model.CasbinRule) (casbins []model.CasbinRule, err error) {
	db := GetDB(c, "mysql").(*gorm.DB)
	if t.ID != 0 {
		db = db.Where("id = ?", t.ID)
	}
	if len(t.Ptype) != 0 {
		db = db.Where("ptype = ?", t.Ptype)
	}
	if len(t.V0) != 0 {
		db = db.Where("v0 = ?", t.V0)
	}
	if len(t.V1) != 0 {
		db = db.Where("v1 = ?", t.V1)
	}
	if len(t.V2) != 0 {
		db = db.Where("v2 = ?", t.V2)
	}

	err = db.Find(&casbins).Error
	return
}
func GetCasbinList(c *gin.Context, t model.CasbinRule) (casbins []model.CasbinRule, err error) {
	db := GetDB(c, "mysql").(*gorm.DB)
	if len(t.Ptype) != 0 {
		db = db.Where("ptype = ?", t.Ptype)
	}
	if len(t.V0) != 0 {
		db = db.Where("v0 = ?", t.V0)
	}
	if len(t.V1) != 0 {
		db = db.Where("v1 LIKE ?", "%"+t.V1+"%")
	}
	if len(t.V2) != 0 {
		db = db.Where("v2 = ?", t.V2)
	}

	err = db.Find(&casbins).Error
	return
}

func DelCasbin(c *gin.Context, t model.CasbinRule) error {
	db := GetDB(c, "mysql").(*gorm.DB)
	return db.Delete(&t).Error
}

func SetCasbin(c *gin.Context, t model.CasbinRule) error {
	db := GetDB(c, "mysql").(*gorm.DB)
	return db.Save(&t).Error
}
