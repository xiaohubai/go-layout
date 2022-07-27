package response

type Menus struct {
	Menu
	Temp
}

type Temp struct {
	Path     string `json:"path"`
	Redirect string `json:"redirect"`
}

type Menu struct {
	ID        int    `json:"ID"`
	CreatedAt string `json:"CreatedAt"`
	UpdatedAt string `json:"UpdatedAt"`
	ParentID  string `json:"parentId"`
	Path      string `json:"path"`
	Name      string `json:"name"`
	Hidden    bool   `json:"hidden"`
	Component string `json:"component"`
	Sort      int    `json:"sort"`
	Meta      Meta   `json:"meta"`
	Children  []Menu `json:"children"`
}

type Meta struct {
	KeepAlive   bool   `json:"keepAlive"`
	DefaultMenu bool   `json:"defaultMenu"`
	Title       string `json:"title"`
	Icon        string `json:"icon"`
	CloseTab    bool   `json:"closeTab"`
}
