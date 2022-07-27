package utils

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/response"
)

// CreateToken 生成token
func createToken(claims model.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(global.Cfg.JWT.SigningKey))
}

// SetToken 生成token
func SetToken(c *gin.Context, user *model.User) (*response.TokenResp, error) {
	claims := model.Claims{
		UID:        user.UID,
		UserName:   user.UserName,
		Phone:      user.Phone,
		RoleID:     user.RoleID,
		RoleName:   user.RoleName,
		State:      user.State,
		BufferTime: global.Cfg.JWT.BufferTime,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 1000,                       // 签名生效时间
			ExpiresAt: time.Now().Unix() + global.Cfg.JWT.ExpiresTime, // 过期时间
			Issuer:    "xiaohubai@outlook.com",                        // 签名的发行者
		},
	}
	token, err := createToken(claims)
	if err != nil {
		return nil, err
	}
	resp := &response.TokenResp{
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}
	return resp, nil
}
