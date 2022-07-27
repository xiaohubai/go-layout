package tables

import (
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"gorm.io/gorm"
)

// Init 注册表结构、导入数据
func Init() {
	err := global.Db.AutoMigrate(
		model.User{},
		model.Menu{},
	)
	if err != nil {
		panic(err)
	}
	if err := InitDBData(); err != nil {
		panic(err)
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
				UID:         "092a82778886abb2",
				UserName:    "admin",
				NickName:    "walle",
				Password:    "891588d2b267551e4609e43cd1159c90",
				Phone:       "13269110806",
				Birth:       "2021-01-01",
				State:       "0",
				RoleID:      "0",
				Salt:        "B`qw&!",
				RoleName:    "管理员",
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
				ParentID:    0,
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
