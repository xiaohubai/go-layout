package model

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	ParentId    int64  `gorm:"column:parent_id;default:0;NOT NULL" json:"parent_id"` // 父节点ID
	Path        string `gorm:"column:path;NOT NULL" json:"path"`                     // 路由path
	Name        string `gorm:"column:name;NOT NULL" json:"name"`                     // 路由名称
	Hidden      int    `gorm:"column:hidden;default:0;NOT NULL" json:"hidden"`       // 0：隐藏，1：展示
	Component   string `gorm:"column:component;NOT NULL" json:"component"`           // 对应前端文件路径
	Title       string `gorm:"column:title;NOT NULL" json:"title"`                   // 显示名称
	Icon        string `gorm:"column:icon;NOT NULL" json:"icon"`                     // 图标
	Sort        int    `gorm:"column:sort;default:1;NOT NULL" json:"sort"`           // 排序标记
	Children    []Menu `gorm:"-" json:"children" `
	CreatedUser string `gorm:"column:create_user;NOT NULL" json:"created_user"` // 创建人
	UpdatedUser string `gorm:"column:update_user;NOT NULL" json:"updated_user"` // 修改人
}
