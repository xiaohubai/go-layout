package dao

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/model"

	"gorm.io/gorm"
)

func SelectUser(c *gin.Context, u *model.User) (user []model.User, count int64, err error) {
	db := GetDB(c, "mysql").(*gorm.DB)
	db = db.Model(model.User{})
	if len(u.Uid) != 0 {
		db = db.Where("uid = ?", u.Uid)
	}
	if len(u.Username) != 0 {
		db = db.Where("username = ?", u.Username)
	}
	if len(u.Password) != 0 {
		db = db.Where("password = ?", u.Password)
	}
	if len(u.Phone) != 0 {
		db = db.Where("phone_number = ?", u.Phone)
	}
	if len(u.State) != 0 {
		db = db.Where("state = ?", u.State)
	}
	if len(u.RoleId) != 0 {
		db = db.Where("role_id = ?", u.RoleId)
	}
	if len(u.RoleName) != 0 {
		db = db.Where("role_name = ?", u.RoleName)
	}

	err = db.Find(&user).Error
	db.Count(&count)
	return
}

func CreateOneUser(c *gin.Context, u []model.User) (err error) {
	db := GetDB(c, "mysql").(*gorm.DB)
	return db.Create(&u).Error
}
