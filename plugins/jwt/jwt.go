package jwt

import (
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"

	"github.com/golang-jwt/jwt"
)

// ParseToken 解析token
func ParseToken(tokenString string) (*model.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &model.Claims{}, func(token *jwt.Token) (i interface{}, e error) {
		return []byte(global.Cfg.JWT.SigningKey), nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, consts.TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, consts.TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, consts.TokenNotValidYet
			} else {
				return nil, consts.TokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*model.Claims); ok && token.Valid {
			return claims, nil
		}
	}
	return nil, consts.TokenInvalid
}
