package service

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/dao"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
)

func AddCasbin(c *gin.Context, r request.CasbinReq) error {
	t := model.CasbinRule{
		Ptype: "p",
		V0:    r.RoleID,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	err := dao.AddCasbin(c, []model.CasbinRule{t})
	if err != nil {
		return err
	}
	return nil
}

func GetCasbinList(c *gin.Context, r request.CasbinListReq) (casbin []response.CasbinResp, err error) {
	t := model.CasbinRule{
		Ptype: r.Ptype,
		V0:    r.RoleID,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	casbins, err := dao.GetCasbinList(c, t)
	if err != nil {
		return nil, err
	}
	for _, v := range casbins {
		data := response.CasbinResp{
			ID:     v.ID,
			RoleID: v.V0,
			Path:   v.V1,
			Method: v.V2,
		}
		casbin = append(casbin, data)
	}
	return
}
