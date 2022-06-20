package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     interface{} `json:"msg"`
	TraceId string      `json:"trace_id"`
}

const (
	CommonOK          = 0    // 成功
	CommonFail        = 4000 // 系统内部错误,未知失败
	CaptchaFail       = 4001 // 验证码获取失败
	ParamsFail        = 4002 // 参数校验错误
	LoginFail         = 4003 // 登录失败
	TokenFail         = 4004 // token无效
	TokenExpired      = 4005 // token授权已过期
	CasbinFail        = 4006 // 权限不足
	CaptchaVerifyFail = 4007 // 验证码校验失败
	RegisterFail      = 4008 // 注册失败:用户已注册
	MenuListFail      = 4009 // 获取路由菜单失败
	CasbinAddFail     = 4010 // 权限添加失败
	CasbinDelFail     = 4011 // 权限删除失败
	CasbinUpdateFail  = 4012 // 权限更新失败
	CasbinListFail    = 4013 // 权限列表失败
	ErrRateLimited    = 4014 // 超出请求频率限制
)

var codeMsg = map[int]string{
	CommonOK:          "成功",
	CommonFail:        "系统内部错误，请稍后再试",
	CaptchaFail:       "验证码获取失败",
	ParamsFail:        "参数校验错误",
	TokenFail:         "获取token失败",
	LoginFail:         "登录失败",
	TokenExpired:      "token授权已过期",
	CasbinFail:        "权限不足",
	RegisterFail:      "注册失败",
	CaptchaVerifyFail: "验证码校验失败",
	MenuListFail:      "获取路由菜单失败",
	CasbinAddFail:     "权限添加失败",
	CasbinDelFail:     "权限删除失败",
	CasbinUpdateFail:  "权限更新失败",
	CasbinListFail:    "权限列表失败",
	ErrRateLimited:    "超出请求频率限制",
}

func Result(c *gin.Context, code int, data, msg interface{}) {
	tracId := ""
	if id, ok := c.Get("X-Trace-ID"); ok {
		tracId = id.(string)
	}
	resp := Resp{
		Code:    code,
		Data:    data,
		Msg:     codeMsg[code],
		TraceId: tracId,
	}
	if e, ok := msg.(error); ok {
		resp.Msg = fmt.Sprintf("%s：%s", resp.Msg, e.Error())
	}
	respBody, _ := json.Marshal(resp)
	c.Set("resp_body", string(respBody))
	c.JSON(http.StatusOK, resp)
}

func Ok(c *gin.Context, data interface{}) {
	if data == nil {
		data = make(map[string]struct{})
	}
	Result(c, CommonOK, data, "")
}

func Fail(c *gin.Context, code int, err interface{}) {
	data := make(map[string]struct{})
	Result(c, code, data, err)
}
