package response

type CasbinResp struct {
	ID     uint   `json:"id"`
	RoleID string `json:"roleID"`
	Path   string `json:"path"`
	Method string `json:"method"`
}
