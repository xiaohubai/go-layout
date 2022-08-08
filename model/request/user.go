package request

type Register struct {
	UserName string `json:"userName"  binding:"required,min=1"`           // 用户名
	NickName string `json:"nickName"`                                     // 昵称
	Password string `json:"password"  binding:"required,min=6,max=20"`    // 密码
	Phone    string `json:"phone" binding:"required,len=11"`              // 手机号
	RoleID   string `json:"roleId" binding:"required,oneof='1' '2'"`      // 角色：0 管理员，1 正常用户，只能输入1
	RoleName string `json:"roleName" binding:"required"`                  // 角色名称
	Birth    string `json:"birth" binding:"required,datetime=2006-01-02"` // 出生日期
	Email    string `json:"email"`                                        // 邮箱地址
	State    string `json:"state"`                                        // 用户状态:(0:初始,1:使用,2:未知,3:停用，4:删除)
}

type LoginReq struct {
	UserName  string `json:"userName" binding:"required,min=1"`        // 用户名
	Password  string `json:"password" binding:"required,min=6,max=20"` // 密码
	Captcha   string `json:"captcha" binding:"required,len=6"`         // 验证码
	CaptchaID string `json:"captchaId" binding:"required"`             // 验证码ID
}

type UserInfoReq struct {
	ID            uint   `json:"id"`
	UID           string `json:"uid"`           // uid
	UserName      string `json:"userName"`      // 用户名：唯一
	Password      string `json:"password"`      // 密码
	NickName      string `json:"nickName"`      // 昵称
	Birth         string `json:"birth"`         // 出生日期:时间戳
	Avatar        string `json:"avatar"`        // 头像:ip+uid/avatar/+值
	RoleID        string `json:"roleId"`        // 角色id：0:管理员，1：正常用户
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
