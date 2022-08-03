package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/model/response"
)

func GetRoleMenus(c *gin.Context) {

	resp := map[string][]response.Menu{
		"menus": {
			{
				Path:      "/layout",
				Name:      "layout",
				Component: "views/layout/index.vue",
				Meta: response.Meta{
					Title: "底层layout",
				},
				Children: []response.Menu{
					{
						ID:        2,
						CreatedAt: "2022-07-09T19:02:48.587+08:00",
						UpdatedAt: "2022-07-09T19:02:48.587+08:00",
						ParentID:  "0",
						Path:      "dashboard",
						Name:      "dashboard",
						Hidden:    false,
						Component: "views/dashboard/index.vue",
						Sort:      0,
						Meta: response.Meta{
							KeepAlive:   false,
							DefaultMenu: false,
							Title:       "仪表盘",
							Icon:        "info-filled",
							CloseTab:    false,
						},
					},
					{
						ID:        3,
						CreatedAt: "2022-07-09T19:02:48.587+08:00",
						UpdatedAt: "2022-07-09T19:02:48.587+08:00",
						ParentID:  "0",
						Path:      "userInfo",
						Name:      "userInfo",
						Hidden:    false,
						Component: "views/userInfo/index.vue",
						Sort:      0,
						Meta: response.Meta{
							KeepAlive:   false,
							DefaultMenu: false,
							Title:       "用户信息",
							Icon:        "info-filled",
							CloseTab:    false,
						},
					},
					{
						ID:        4,
						Path:      "404",
						Name:      "404",
						Component: "views/error/index.vue",
						Hidden:    true,
						Meta: response.Meta{
							Title: "404",
						},
					},
					{
						ID:        5,
						CreatedAt: "2022-07-09T19:02:48.587+08:00",
						UpdatedAt: "2022-07-09T19:02:48.587+08:00",
						ParentID:  "0",
						Path:      "person",
						Name:      "person",
						Hidden:    true,
						Component: "views/person/index.vue",
						Sort:      0,
						Meta: response.Meta{
							KeepAlive:   false,
							DefaultMenu: false,
							Title:       "个人信息",
							Icon:        "info-filled",
							CloseTab:    false,
						},
					},
					{
						ID:        6,
						CreatedAt: "2022-07-09T19:02:48.587+08:00",
						UpdatedAt: "2022-07-09T19:02:48.587+08:00",
						ParentID:  "0",
						Path:      "casbin",
						Name:      "casbin",
						Hidden:    false,
						Component: "views/casbin/index.vue",
						Sort:      0,
						Meta: response.Meta{
							KeepAlive:   false,
							DefaultMenu: false,
							Title:       "权限表",
							Icon:        "info-filled",
							CloseTab:    false,
						},
					},
				},
			},
		},
	}

	response.Ok(c, resp)
}
