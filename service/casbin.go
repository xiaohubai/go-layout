package service

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/dao"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
)

// AddCasbin 增加权限路由
func AddCasbin(c *gin.Context, r request.CasbinReq) error {
	t := model.CasbinRule{
		Ptype: "p",
		V0:    r.RoleID,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	casbins, err := dao.CasbinList(c, t)
	if err != nil {
		return err
	}
	if len(casbins) > 0 {
		return fmt.Errorf("该权限路由已存在")
	}
	err = dao.AddCasbin(c, []model.CasbinRule{t})
	if err != nil {
		return err
	}
	return nil
}

// GetCasbinList 获取全部权限路由列表
func GetCasbinList(c *gin.Context, r request.CasbinListReq) (resp *response.PageResp, err error) {
	t := model.CasbinRule{
		Ptype: r.Ptype,
		V0:    r.RoleID,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	pageInfo := response.PageInfo{
		Page:     r.Page,
		PageSize: r.PageSize,
	}
	casbins, total, err := dao.GetCasbinList(c, t, pageInfo)
	if err != nil {
		return nil, err
	}

	var casbin []response.CasbinResp
	for _, v := range casbins {
		data := response.CasbinResp{
			ID:     v.ID,
			RoleID: v.V0,
			Path:   v.V1,
			Method: v.V2,
		}
		casbin = append(casbin, data)
	}

	resp = &response.PageResp{
		List:     casbin,
		Total:    total,
		Page:     r.Page,
		PageSize: r.PageSize,
	}
	return
}

// DelCasbin 删除一条权限路由
func DelCasbin(c *gin.Context, r request.DelCasbinReq) error {
	t := model.CasbinRule{
		ID: uint(r.ID),
	}
	casbins, err := dao.CasbinList(c, t)
	if err != nil {
		return err
	}
	if len(casbins) <= 0 {
		return fmt.Errorf("该权限路由不存在")
	}
	err = dao.DelCasbin(c, t)
	if err != nil {
		return err
	}
	return nil
}

// SetCasbin 更新权限路由
func SetCasbin(c *gin.Context, r request.SetCasbinReq) error {
	t := model.CasbinRule{
		ID:    uint(r.ID),
		Ptype: r.Ptype,
		V0:    r.RoleID,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	casbins, err := dao.CasbinList(c, model.CasbinRule{ID: uint(r.ID)})
	if err != nil {
		return err
	}
	if len(casbins) <= 0 {
		return fmt.Errorf("该权限路由不存在")
	}
	err = dao.SetCasbin(c, t)
	if err != nil {
		return err
	}
	return nil
}
