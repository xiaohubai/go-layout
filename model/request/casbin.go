package request

type CasbinReq struct {
	RoleID string `json:"roleId"  binding:"required"`
	Path   string `json:"path" binding:"required"`
	Method string `json:"method" binding:"required"`
}

type CasbinListReq struct {
	Ptype  string `json:"ptype"`
	RoleID string `json:"roleId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

type DelCasbinReq struct {
	ID int `json:"id"  binding:"required"`
}

type SetCasbinReq struct {
	ID     int    `json:"id"  binding:"required"`
	Ptype  string `json:"ptype"`
	RoleID string `json:"roleId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
