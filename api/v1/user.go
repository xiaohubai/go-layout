package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

func Register(c *gin.Context) {
	var r request.Register
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFail, err)
		return
	}
	u := &model.User{
		UserName: r.UserName,
		Password: r.Password,
		Phone:    r.Phone,
		RoleID:   r.RoleID,
		RoleName: r.RoleName,
		Birth:    r.Birth,
	}

	err := service.Register(c, u)
	if err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "注册失败", err)))
		response.Fail(c, response.RegisterFail, err)
		return
	}
	response.Ok(c, nil)
}

func Login(c *gin.Context) {
	var r request.LoginReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFail, err)
		return
	}
	if !store.Verify(r.CaptchaID, r.Captcha, true) {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", "验证码校验失败"))
		response.Fail(c, response.CaptchaVerifyFail, nil)
		return
	}
	u := &model.User{UserName: r.UserName, Password: r.Password}
	if loginResp, err := service.Login(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.LoginFail, err)
	} else {
		response.Ok(c, loginResp)
	}
}

func UserInfo(c *gin.Context) {
	claims, ok := c.Get("claims")
	if !ok {
		response.Fail(c, response.TokenFail, nil)
		return
	}
	r := claims.(*model.Claims)
	u := &model.User{
		UID: r.UID,
	}
	if userInfoResp, err := service.UserInfoList(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.GetUserInfoFaild, err)
	} else {

		response.Ok(c, map[string]interface{}{"userInfo": userInfoResp[0]})
	}
}

func UserInfoList(c *gin.Context) {
	var r request.UserInfoReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFail, err)
		return
	}

	u := &model.User{
		UserName:    r.UserName,
		UID:         r.UID,
		RoleID:      r.RoleID,
		RoleName:    r.RoleName,
		Phone:       r.Phone,
		Wechat:      r.Wechat,
		State:       r.State,
		CreatedUser: r.CreatedUser,
		UpdatedUser: r.UpdatedUser,
	}
	if userInfoResp, err := service.UserInfoList(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.LoginFail, err)
	} else {
		response.Ok(c, userInfoResp)
	}
}

func SetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*model.Claims)
	var r request.UserInfoReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFail, err)
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
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.LoginFail, err)
	} else {
		response.Ok(c, nil)
	}
}
