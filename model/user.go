package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UID           string `gorm:"column:uid;NOT NULL" json:"uid"`                                                              // uid
	UserName      string `gorm:"column:username;NOT NULL" json:"username"`                                                    // 用户名：唯一
	NickName      string `gorm:"column:nickname;NOT NULL" json:"nickname"`                                                    // 昵称
	Password      string `gorm:"column:password;NOT NULL" json:"password"`                                                    // md5密码
	Salt          string `gorm:"column:salt;NOT NULL" json:"salt"`                                                            // 加盐
	Birth         string `gorm:"column:birth;NOT NULL" json:"birth"`                                                          // 出生日期
	Avatar        string `gorm:"column:avatar;default:https://qmplusimg.henrongyi.top/gva_header.jpg;NOT NULL" json:"avatar"` // 头像:ip+uid/avatar/+值
	RoleID        string `gorm:"column:role_id;NOT NULL" json:"role_id"`                                                      // 角色id：0:管理员，1：正常用户
	RoleName      string `gorm:"column:role_name;NOT NULL" json:"role_name"`                                                  // 角色名称
	Phone         string `gorm:"column:phone;NOT NULL" json:"phone"`                                                          // 手机号：唯一
	Wechat        string `gorm:"column:wechat;NOT NULL" json:"wechat"`                                                        // 微信号
	Email         string `gorm:"column:email;NOT NULL" json:"email"`                                                          // 邮箱地址
	State         string `gorm:"column:state;default:1;NOT NULL" json:"state"`                                                // 用户状态:(0:初始,1:使用,2:未知,3:停用，4:删除)
	DefaultRouter string `gorm:"column:default_router;default:about;NOT NULL" json:"default_router"`                          // 默认路由
	SideMode      string `gorm:"column:side_mode;default:dark;NOT NULL" json:"side_mode"`                                     // layout边框颜色
	BaseColor     string `gorm:"column:base_color;default:#fff;NOT NULL" json:"base_color"`                                   // layout基础颜色
	ActiveColor   string `gorm:"column:active_color;default:#1890ff;NOT NULL" json:"active_color"`                            // 选中颜色
	CreatedUser   string `gorm:"column:create_user;NOT NULL" json:"created_user"`                                             // 创建人
	UpdatedUser   string `gorm:"column:update_user;NOT NULL" json:"updated_user"`                                             // 修改人
}
