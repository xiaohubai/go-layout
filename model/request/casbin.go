package request

type CasbinReq struct {
	RoleId string `json:"roleId"  binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
