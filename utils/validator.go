package utils

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
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
