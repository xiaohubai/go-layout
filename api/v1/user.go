package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"

	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/metrics"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

func Register(c *gin.Context) {
	var r request.Register
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}
	u := &model.User{
		UserName: r.UserName,
		NickName: r.NickName,
		Password: r.Password,
		Phone:    r.Phone,
		RoleID:   r.RoleID,
		RoleName: r.RoleName,
		Birth:    r.Birth,
		State:    r.State,
		Email:    r.Email,
	}

	err := service.Register(c, u)
	if err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "注册失败", err)))
		response.Fail(c, response.RegisterFailed, err)
		return
	}
	response.Ok(c, nil)
}

func Login(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	metrics.CounterIncM("login")

	var r request.LoginReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFailed, err)
		return
	}
	if !store.Verify(r.CaptchaID, r.Captcha, true) {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", "验证码校验失败"))
		response.Fail(c, response.CaptchaVerifyFailed, nil)
		return
	}
	u := &model.User{UserName: r.UserName, Password: r.Password}
	if loginResp, err := service.Login(c, u); err != nil {
		span.LogFields(log.Object("service.Login(c, u)", u), log.Object("error", err))
		response.Fail(c, response.LoginFailed, err)
	} else {
		response.Ok(c, loginResp)
	}
}

//UserInfo 通过token的UID获取用户信息
func UserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*model.Claims)
	u := &model.User{
		UID: claims.UID,
	}
	if userInfoResp, err := service.UserInfoList(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "获取用户信息失败", err)))
		response.Fail(c, response.GetUserInfoFailed, err)
	} else {
		response.Ok(c, map[string]interface{}{"userInfo": userInfoResp[0]})
	}
}

func UserInfoList(c *gin.Context) {
	claims := c.MustGet("claims").(*model.Claims)
	if claims.RoleID != consts.AdminRoleID {
		response.Fail(c, response.GetUserInfoFailed, fmt.Errorf("非超级管理员"))
		return
	}
	if userInfoResp, err := service.UserInfoList(c, &model.User{}); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "获取用户列表失败", err)))
		response.Fail(c, response.GetUserInfoFailed, err)
	} else {
		response.Ok(c, userInfoResp)
	}
}

func SetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*model.Claims)
	var r request.UserInfoReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFailed, err)
		return
	}

	u := &model.User{
		UID:         r.UID,
		RoleID:      r.RoleID,
		RoleName:    r.RoleName,
		Phone:       r.Phone,
		Wechat:      r.Wechat,
		State:       r.State,
		Email:       r.Email,
		SideMode:    r.SideMode,
		Password:    r.Password,
		Avatar:      r.Avatar,
		UpdatedUser: claims.UserName,
	}
	if err := service.SetUserInfo(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "设置用户信息失败", err)))
		response.Fail(c, response.SetUserInfoFailed, err)
	} else {
		response.Ok(c, nil)
	}
}
