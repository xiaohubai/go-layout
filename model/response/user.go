package response

type TokenResp struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}

type LoginResp struct {
	UserInfo  UserInfoResp `json:"userInfo"`
	TokenInfo TokenResp    `json:"tokenInfo"`
}

type UserInfoResp struct {
	ID            uint   `json:"id"`
	UID           string `json:"uid"`           // uid
	UserName      string `json:"userName"`      // 用户名：唯一
	NickName      string `json:"nickName"`      // 昵称
	Birth         string `json:"birth"`         // 出生日期:时间戳
	Avatar        string `json:"avatar"`        // 头像:ip+uid/avatar/+值
	RoleID        string `json:"roleID"`        // 角色id：0:管理员，1：正常用户
	RoleName      string `json:"roleName"`      // 角色名称
	Phone         string `json:"phone"`         // 手机号：唯一
	Wechat        string `json:"wechat"`        // 微信号
	Email         string `json:"email"`         // 邮箱地址
	State         string `json:"state"`         // 用户状态:(0:初始,1:使用,2:未知,3:停用，4:删除)
	DefaultRouter string `json:"defaultRouter"` // 默认路由
	SideMode      string `json:"sideMode"`      // layout边框颜色
	BaseColor     string `json:"baseColor"`     // layout基础颜色
	ActiveColor   string `json:"activeColor"`   // 选中颜色
	CreatedUser   string `json:"createdUser"`   // 创建人
	UpdatedUser   string `json:"updatedUser"`   // 修改人
	CreateAt      string `json:"createAt"`      // 创建时间"
	UpdateAt      string `json:"updateAt"`      // 更新时间
}
