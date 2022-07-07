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
		Username: r.Username,
		Password: r.Password,
		Phone:    r.Phone,
		RoleId:   r.RoleId,
		RoleName: r.RoleName,
		Birth:    utils.StrToTime(r.Birth, "2006-01-02"),
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
	if !store.Verify(r.CaptchaId, r.Captcha, true) {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", "验证码校验失败"))
		response.Fail(c, response.CaptchaVerifyFail, nil)
		return
	}
	u := &model.User{Username: r.Username, Password: r.Password}
	if loginResp, err := service.Login(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.LoginFail, err)
	} else {
		response.Ok(c, loginResp)
	}
}

func UserInfo(c *gin.Context) {
	var r request.UserInfoReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "参数校验错误", err)))
		response.Fail(c, response.ParamsFail, err)
		return
	}

	u := &model.User{
		Username:    r.Username,
		Uid:         r.Uid,
		RoleId:      r.RoleId,
		RoleName:    r.RoleName,
		Phone:       r.Phone,
		Wechat:      r.Wechat,
		State:       r.State,
		CreatedUser: r.CreatedUser,
		UpdatedUser: r.UpdatedUser,
	}
	if userInfoResp, err := service.UserInfo(c, u); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "登录失败", err)))
		response.Fail(c, response.LoginFail, err)
	} else {
		response.Ok(c, userInfoResp)
	}
}
