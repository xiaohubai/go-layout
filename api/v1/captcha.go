package v1

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"

	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/utils"
	"go.uber.org/zap"
)

var store = base64Captcha.DefaultMemStore

//Captcha 生成验证码
func Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.Cfg.Captcha.ImgHeight, global.Cfg.Captcha.ImgWidth,
		global.Cfg.Captcha.KeyLong, 0.7, 80)
	cp := base64Captcha.NewCaptcha(driver, store)
	if id, b64s, err := cp.Generate(); err != nil {
		global.Log.Error(utils.TraceId(c), zap.Any("key", "func"), zap.Any("msg", fmt.Sprintf("%s:%s", "验证码获取失败", err)))
		response.Fail(c, response.CaptchaFail, nil)
	} else {
		response.Ok(c, response.CaptchaResp{CaptchaID: id, PicPath: b64s, CaptchaLength: global.Cfg.Captcha.KeyLong})
	}
}
