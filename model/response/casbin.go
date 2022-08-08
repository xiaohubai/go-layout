package response

type CasbinResp struct {
	ID     uint   `json:"id"`
	RoleID string `json:"roleId"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
