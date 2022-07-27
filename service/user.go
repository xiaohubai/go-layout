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
	userInfo, count, err := dao.SelectUser(c, &model.User{UserName: u.UserName})
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
		ID:            userInfo[0].Model.ID,
		UID:           userInfo[0].UID,
		UserName:      userInfo[0].UserName,
		NickName:      userInfo[0].NickName,
		Birth:         userInfo[0].Birth,
		Avatar:        userInfo[0].Avatar,
		RoleID:        userInfo[0].RoleID,
		RoleName:      userInfo[0].RoleName,
		Phone:         userInfo[0].Phone,
		Wechat:        userInfo[0].Wechat,
		Email:         userInfo[0].Email,
		DefaultRouter: userInfo[0].DefaultRouter,
		State:         userInfo[0].State,
		SideMode:      userInfo[0].SideMode,
		BaseColor:     userInfo[0].BaseColor,
		ActiveColor:   userInfo[0].ActiveColor,
		CreatedUser:   userInfo[0].CreatedUser,
		UpdatedUser:   userInfo[0].UpdatedUser,
		CreateAt:      utils.TimeToStr(userInfo[0].Model.CreatedAt, "2006-01-02 15:04:05"),
		UpdateAt:      utils.TimeToStr(userInfo[0].Model.UpdatedAt, "2006-01-02 15:04:05"),
	}

	tokenInfo, err := utils.SetToken(c, &userInfo[0])
	resp := &response.LoginResp{
		UserInfo:  user,
		TokenInfo: *tokenInfo,
	}
	return resp, err
}

func UserInfoList(c *gin.Context, u *model.User) ([]response.UserInfoResp, error) {
	userInfos, _, err := dao.SelectUser(c, u)
	if err != nil {
		return nil, fmt.Errorf("系统内部错误")
	}
	resp := make([]response.UserInfoResp, 0)
	for _, v := range userInfos {
		user := response.UserInfoResp{
			ID:            v.Model.ID,
			UID:           v.UID,
			UserName:      v.UserName,
			NickName:      v.NickName,
			Birth:         v.Birth,
			Avatar:        v.Avatar,
			RoleID:        v.RoleID,
			RoleName:      v.RoleName,
			Phone:         v.Phone,
			Wechat:        v.Wechat,
			Email:         v.Email,
			State:         v.State,
			CreatedUser:   v.CreatedUser,
			UpdatedUser:   v.UpdatedUser,
			DefaultRouter: v.DefaultRouter,
			SideMode:      v.SideMode,
			BaseColor:     v.BaseColor,
			ActiveColor:   v.ActiveColor,
			CreateAt:      utils.TimeToStr(v.Model.CreatedAt, "2006-01-02 15:04:05"),
			UpdateAt:      utils.TimeToStr(v.Model.UpdatedAt, "2006-01-02 15:04:05"),
		}
		resp = append(resp, user)
	}
	return resp, err
}

func SetUserInfo(c *gin.Context, u *model.User) error {
	err := dao.UpdateUser(c, u)
	fmt.Println(err)
	return err
}
func Register(c *gin.Context, u *model.User) (err error) {
	if _, err = predis.Get(c.Request.Context(), utils.Md5([]byte(u.UserName))); err == nil {
		return fmt.Errorf("用户名已存在")
	}
	if err != redis.Nil {
		return fmt.Errorf("系统内部错误")
	}

	user := &model.User{UserName: u.UserName}
	_, count, err := dao.SelectUser(c, user)
	if err != nil {
		return fmt.Errorf("系统内部错误")
	}
	if count >= 1 {
		return fmt.Errorf("用户名已存在")
	}
	if err := predis.Set(c.Request.Context(), u.UserName, "1", 5*time.Minute); err != nil {
		return fmt.Errorf("系统内部错误")
	}

	u.CreatedUser = u.UserName
	u.UpdatedUser = u.UserName
	u.UID = utils.TraceId(c)
	u.Salt = utils.RandString(7)
	u.Avatar = "avatar.jpg"
	u.NickName = "hi"
	u.State = "1"
	u.Password = utils.Md5([]byte(u.Password + u.Salt))

	err = dao.CreateOneUser(c, []model.User{*u})
	_ = predis.Del(c.Request.Context(), utils.Md5([]byte(u.UserName)))
	return err
}
