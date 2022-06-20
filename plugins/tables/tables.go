package tables

import (
	"fmt"

	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/utils"
	"gorm.io/gorm"
)

// Init 注册表结构、导入数据
func Init() {
	err := global.Db.AutoMigrate(
		model.User{},
		model.Menu{},
	)
	if err != nil {
		panic(fmt.Errorf("tables register failed: %s \n", err))
	}
	if err := InitDBData(); err != nil {
		panic(fmt.Errorf("tables init insert failed: %s \n", err))
	}
}

var goFuncs = []func() error{
	CreateUser,
	CreateMenu,
}

func InitDBData() error {
	for _, f := range goFuncs {
		err := f()
		if err != nil {
			return err
		}
	}
	return nil
}

func CreateUser() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.User{}).RowsAffected == 1 {
			return nil
		}

		user := []model.User{
			{
				Model:       gorm.Model{ID: 1},
				Uid:         utils.UUID(),
				Username:    "admin",
				Password:    "891588d2b267551e4609e43cd1159c90",
				Phone:       "13269110806",
				Birth:       utils.StrToTime("2021-01-01", "2006-01-02"),
				State:       "0",
				RoleId:      "0",
				Salt:        "B`qw&!",
				RoleName:    "管理员",
				Avatar:      "avatar.jpg",
				CreatedUser: "admin",
				UpdatedUser: "admin",
			},
		}
		if err := tx.Create(&user).Error; err != nil {
			return err
		}
		return nil
	})
}

func CreateMenu() error {
	return global.Db.Transaction(func(tx *gorm.DB) error {
		if tx.Where("id IN ?", []int{1}).Find(&[]model.Menu{}).RowsAffected == 1 {
			return nil
		}
		menus := []model.Menu{
			{
				Model:       gorm.Model{ID: 1},
				ParentId:    0,
				Path:        "dashboard",
				Name:        "dashboard",
				Hidden:      0,
				Component:   "view/dashboard/index.vue",
				Title:       "仪表盘",
				Icon:        "setting",
				Sort:        1,
				CreatedUser: "admin",
				UpdatedUser: "admin",
			},
		}
		if err := tx.Create(&menus).Error; err != nil {
			return err
		}
		return nil
	})
}
