package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

func AddCasbin(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	var r request.CasbinReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		span.LogFields(log.Object("utils.ShouldBindJSON(c, &r)", r), log.Object("error", err))
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if err := service.AddCasbin(c, r); err != nil {
		span.LogFields(log.Object("service.AddCasbin(c, r)", r), log.Object("error", err))
		response.Fail(c, response.CasbinAddFailed, nil)
	} else {
		response.Ok(c, nil)
	}

}

func GetCasbinList(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	claims := c.MustGet("claims").(*model.Claims)
	if claims.RoleID != consts.AdminRoleID {
		response.Fail(c, response.GetUserInfoFailed, fmt.Errorf("非超级管理员"))
		return
	}
	var r request.CasbinListReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		span.LogFields(log.Object("utils.ShouldBindJSON(c, &r)", r), log.Object("error", err))
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	if casbinListResp, err := service.GetCasbinList(c, r); err != nil {
		span.LogFields(log.Object("service.GetCasbinList(c)", ""), log.Object("error", err))
		response.Fail(c, response.GetCasbinListFailed, nil)
	} else {
		response.Ok(c, casbinListResp)
	}

}
