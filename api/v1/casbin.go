package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

// AddCasbin 添加权限路由
func AddCasbin(c *gin.Context) {
	var r request.CasbinReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if err := service.AddCasbin(c, r); err != nil {
		response.Fail(c, response.CasbinAddFailed, err)
	} else {
		response.Ok(c, nil)
	}

}

// GetCasbinList 获取权限路由列表
func GetCasbinList(c *gin.Context) {
	if !utils.IsAdminID(c) {
		response.Fail(c, response.NotAdminID, nil)
		return
	}

	var r request.CasbinListReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if casbinListResp, err := service.GetCasbinList(c, r); err != nil {
		response.Fail(c, response.GetCasbinListFailed, nil)
	} else {
		response.Ok(c, casbinListResp)
	}
}

// DelCasbin 删除权限路由
func DelCasbin(c *gin.Context) {
	if !utils.IsAdminID(c) {
		response.Fail(c, response.NotAdminID, nil)
		return
	}
	var r request.DelCasbinReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if err := service.DelCasbin(c, r); err != nil {
		response.Fail(c, response.CasbinDelFailed, err)
	} else {
		response.Ok(c, nil)
	}

}

// SetCasbin 更新权限路由
func SetCasbin(c *gin.Context) {
	if !utils.IsAdminID(c) {
		response.Fail(c, response.NotAdminID, nil)
		return
	}

	var r request.SetCasbinReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if err := service.SetCasbin(c, r); err != nil {
		response.Fail(c, response.SetCasbinFailed, err)
	} else {
		response.Ok(c, nil)
	}

}
