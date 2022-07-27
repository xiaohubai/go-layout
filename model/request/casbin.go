package request

type CasbinReq struct {
	RoleID string `json:"roleID"  binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}
