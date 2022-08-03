package request

type CasbinReq struct {
	RoleID string `json:"roleID"  binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}

type CasbinListReq struct {
	Ptype  string `json:"ptype"`
	RoleID string `json:"roleID"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
