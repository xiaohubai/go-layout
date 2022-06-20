package request

type TokenReq struct {
	Username string `json:"username" binding:"required,min=1"`        // 用户名
	Password string `json:"password" binding:"required,min=6,max=20"` // 密码
}
