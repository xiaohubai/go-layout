package response

type TokenResp struct {
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expires_at"`
}

type LoginResp struct {
	UserInfo  UserInfoResp `json:"user_info"`
	TokenInfo TokenResp    `json:"toekn_info"`
}

type UserInfoResp struct {
	Id          uint   `json:"id"`
	Uid         string `json:"uid"`          // uid
	Username    string `json:"username"`     // 用户名：唯一
	Nick        string `json:"nick"`         // 昵称
	Birth       string `json:"birth"`        // 出生日期:时间戳
	Avatar      string `json:"avatar"`       // 头像:ip+uid/avatar/+值
	RoleId      string `json:"role_id"`      // 角色id：0:管理员，1：正常用户
	RoleName    string `json:"role_name"`    // 角色名称
	Phone       string `json:"phone"`        // 手机号：唯一
	Wechat      string `json:"wechat"`       // 微信号
	Email       string `json:"email"`        // 邮箱地址
	State       string `json:"state"`        // 用户状态:(0:初始,1:使用,2:未知,3:停用，4:删除)
	CreatedUser string `json:"created_user"` // 创建人
	UpdatedUser string `json:"updated_user"` // 修改人
	CreateAt    string `json:"create_at"`    // 创建时间"
	UpdateAt    string `json:"update_at"`    // 更新时间
}
