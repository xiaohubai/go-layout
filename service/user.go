package service

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/xiaohubai/go-layout/dao"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/response"
	predis "github.com/xiaohubai/go-layout/plugins/redis"
	"github.com/xiaohubai/go-layout/utils"
)

func Login(c *gin.Context, u *model.User) (result *response.LoginResp, err error) {
	userInfo, count, err := dao.SelectUser(c, &model.User{Username: u.Username})
	if err != nil {
		return nil, fmt.Errorf("系统内部错误")
	}
	if count != 1 {
		return nil, fmt.Errorf("用户不存在")
	}

	if userInfo[0].Password != utils.Md5([]byte(u.Password+userInfo[0].Salt)) {
		return nil, fmt.Errorf("密码错误")
	}

	user := response.UserInfoResp{
		Id:          userInfo[0].Model.ID,
		Uid:         userInfo[0].Uid,
		Username:    userInfo[0].Username,
		Nick:        userInfo[0].Nick,
		Birth:       utils.TimeToStr(userInfo[0].Birth, "2006-01-02"),
		Avatar:      userInfo[0].Avatar,
		RoleId:      userInfo[0].RoleId,
		RoleName:    userInfo[0].RoleName,
		Phone:       userInfo[0].Phone,
		Wechat:      userInfo[0].Wechat,
		Email:       userInfo[0].Email,
		CreatedUser: userInfo[0].CreatedUser,
		UpdatedUser: userInfo[0].UpdatedUser,
		CreateAt:    utils.TimeToStr(userInfo[0].Model.CreatedAt, "2006-01-02 15:04:05"),
		UpdateAt:    utils.TimeToStr(userInfo[0].Model.UpdatedAt, "2006-01-02 15:04:05"),
	}

	tokenInfo, err := utils.SetToken(c, &userInfo[0])
	resp := &response.LoginResp{
		UserInfo:  user,
		TokenInfo: *tokenInfo,
	}
	return resp, err
}

func UserInfo(c *gin.Context, u *model.User) ([]response.UserInfoResp, error) {
	userInfos, _, err := dao.SelectUser(c, &model.User{Username: u.Username})
	if err != nil {
		return nil, fmt.Errorf("系统内部错误")
	}
	resp := make([]response.UserInfoResp, 0)
	for _, v := range userInfos {
		user := response.UserInfoResp{
			Id:          v.Model.ID,
			Uid:         v.Uid,
			Username:    v.Username,
			Nick:        v.Nick,
			Birth:       utils.TimeToStr(v.Birth, "2006-01-02"),
			Avatar:      v.Avatar,
			RoleId:      v.RoleId,
			RoleName:    v.RoleName,
			Phone:       v.Phone,
			Wechat:      v.Wechat,
			Email:       v.Email,
			CreatedUser: v.CreatedUser,
			UpdatedUser: v.UpdatedUser,
			CreateAt:    utils.TimeToStr(v.Model.CreatedAt, "2006-01-02 15:04:05"),
			UpdateAt:    utils.TimeToStr(v.Model.UpdatedAt, "2006-01-02 15:04:05"),
		}
		resp = append(resp, user)
	}

	return resp, err
}

func Register(c *gin.Context, u *model.User) (err error) {
	if _, err = predis.Get(c.Request.Context(), utils.Md5([]byte(u.Username))); err == nil {
		return fmt.Errorf("用户名已存在")
	}
	if err != redis.Nil {
		return fmt.Errorf("系统内部错误")
	}

	user := &model.User{Username: u.Username}
	_, count, err := dao.SelectUser(c, user)
	if err != nil {
		return fmt.Errorf("系统内部错误")
	}
	if count >= 1 {
		return fmt.Errorf("用户名已存在")
	}
	if err := predis.Set(c.Request.Context(), u.Username, "1", 5*time.Minute); err != nil {
		return fmt.Errorf("系统内部错误")
	}

	u.CreatedUser = u.Username
	u.UpdatedUser = u.Username
	u.Uid = utils.TraceId(c)
	u.Salt = utils.RandString(7)
	u.Avatar = "avatar.jpg"
	u.Nick = "hi"
	u.State = "1"
	u.Password = utils.Md5([]byte(u.Password + u.Salt))

	err = dao.CreateOneUser(c, []model.User{*u})
	_ = predis.Del(c.Request.Context(), utils.Md5([]byte(u.Username)))
	return err
}
