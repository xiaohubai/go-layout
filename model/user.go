package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Uid         string    `gorm:"column:uid;NOT NULL" json:"uid"`                  // uid
	Username    string    `gorm:"column:username;NOT NULL" json:"username"`        // 用户名：唯一
	Nick        string    `gorm:"column:nick;NOT NULL" json:"nick"`                // 昵称
	Password    string    `gorm:"column:password;NOT NULL" json:"password"`        // md5密码
	Salt        string    `gorm:"column:salt;NOT NULL" json:"salt"`                // 加盐
	Birth       time.Time `json:"birth"`                                           // 出生日期:时间戳
	Avatar      string    `gorm:"column:avatar;NOT NULL" json:"avatar"`            // 头像:ip+uid/avatar/+值
	RoleId      string    `gorm:"column:role_id;NOT NULL" json:"role_id"`          // 角色id：0:管理员，1：正常用户
	RoleName    string    `gorm:"column:role_name;NOT NULL" json:"role_name"`      // 角色名称
	Phone       string    `gorm:"column:phone;NOT NULL" json:"phone"`              // 手机号：唯一
	Wechat      string    `gorm:"column:wechat;NOT NULL" json:"wechat"`            // 微信号
	Email       string    `gorm:"column:email;NOT NULL" json:"email"`              // 邮箱地址
	State       string    `gorm:"column:state;default:1;NOT NULL" json:"state"`    // 用户状态:(0:初始,1:使用,2:未知,3:停用，4:删除)
	CreatedUser string    `gorm:"column:create_user;NOT NULL" json:"created_user"` // 创建人
	UpdatedUser string    `gorm:"column:update_user;NOT NULL" json:"updated_user"` // 修改人
}
