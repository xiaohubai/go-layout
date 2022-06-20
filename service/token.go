package service

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/dao"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/utils"
)

func Token(c *gin.Context, u *model.User) (*response.TokenResp, error) {
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
	result, err := utils.SetToken(c, &userInfo[0])
	return result, err
}
