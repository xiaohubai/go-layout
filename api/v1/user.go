package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/kafka"
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
		response.Fail(c, response.RegisterFailed, err)
		return
	}
	response.Ok(c, nil)
}

func Login(c *gin.Context) {
	metrics.CounterIncM("login")

	var r request.LoginReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		response.Fail(c, response.ParamsFailed, err)
		return
	}
	if !store.Verify(r.CaptchaID, r.Captcha, true) {
		response.Fail(c, response.CaptchaVerifyFailed, nil)
		return
	}
	u := &model.User{UserName: r.UserName, Password: r.Password}
	if loginResp, err := service.Login(c, u); err != nil {
		response.Fail(c, response.LoginFailed, err)
	} else {
		//登录成功后，将用户信息和token发送到kafka
		kafka.WriteToKafka(consts.TopicOfLoginInfo, utils.JsonToString(loginResp))
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
		response.Fail(c, response.GetUserInfoFailed, err)
	} else {
		response.Ok(c, userInfoResp)
	}
}

func SetUserInfo(c *gin.Context) {
	claims := c.MustGet("claims").(*model.Claims)
	var r request.UserInfoReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
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
		response.Fail(c, response.SetUserInfoFailed, err)
	} else {
		response.Ok(c, nil)
	}
}
