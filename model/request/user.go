package request

type Register struct {
	Username string `json:"username"  binding:"required,min=1"`           // 用户名
	Password string `json:"password"  binding:"required,min=6,max=20"`    // 密码
	Phone    string `json:"phone" binding:"required,len=11"`              // 手机号
	RoleId   int    `json:"role_id" binding:"required,min=1,max=1"`       // 角色：0 管理员，1 正常用户，只能输入1
	RoleName string `json:"role_name" binding:"required"`                 // 角色名称
	Birth    string `json:"birth" binding:"required,datetime=2006-01-02"` // 出生日期
}

type LoginReq struct {
	Username  string `json:"username" binding:"required,min=1"`        // 用户名
	Password  string `json:"password" binding:"required,min=6,max=20"` // 密码
	Captcha   string `json:"captcha" binding:"required,len=6"`         // 验证码
	CaptchaId string `json:"captcha_id" binding:"required"`            // 验证码ID
}

type UserInfoReq struct {
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
