package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/model"
)

func ShouldBindJSON(c *gin.Context, obj interface{}) error {
	err := c.ShouldBindJSON(obj)
	if err != nil {
		//翻译错误信息
		if errs, ok := err.(validator.ValidationErrors); ok {
			v := c.Value("trans")
			trans, _ := v.(ut.Translator)
			var e string
			for _, v := range errs.Translate(trans) {
				e += fmt.Sprintf("%s ", v)
			}
			return fmt.Errorf(e)
		}
	}
	return err
}

func IsAdminID(c *gin.Context) bool {
	claims := c.MustGet("claims").(*model.Claims)
	return claims.RoleID == consts.AdminRoleID
}
