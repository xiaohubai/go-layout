package response

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Resp struct {
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Msg     interface{} `json:"msg"`
	TraceId string      `json:"traceID"`
}

const (
	CommonOK            = 0    // 成功
	CommonFailed        = 4000 // 系统内部错误
	CaptchaFailed       = 4001 // 验证码获取失败
	ParamsFailed        = 4002 // 参数校验错误
	LoginFailed         = 4003 // 登录失败
	TokenFailed         = 4004 // token无效
	TokenExpired        = 4005 // token授权已过期
	CasbinFailed        = 4006 // 权限不足
	CaptchaVerifyFailed = 4007 // 验证码校验失败
	RegisterFailed      = 4008 // 注册失败:用户已注册
	MenuListFailed      = 4009 // 获取路由菜单失败
	CasbinAddFailed     = 4010 // 权限添加失败
	CasbinDelFailed     = 4011 // 权限删除失败
	CasbinUpdateFailed  = 4012 // 权限更新失败
	CasbinListFailed    = 4013 // 权限列表失败
	RateLimited         = 4014 // 超出请求频率限制
	FileWithExcelFailed = 4015 // 文件不是excel
	FileReportFailed    = 4016 // 文件上传失败
	FileOpenFailed      = 4017 // 文件打开失败
	GetUserInfoFailed   = 4018 // 获取用户信息失败
	SetUserInfoFailed   = 4019 // 更新用户信息失败
	GetCasbinListFailed = 4020 // 获取权限表信息失败
)

var codeMsg = map[int]string{
	CommonOK:            "成功",
	CommonFailed:        "系统内部错误",
	CaptchaFailed:       "验证码获取失败",
	ParamsFailed:        "参数校验错误",
	TokenFailed:         "获取token失败",
	LoginFailed:         "登录失败",
	TokenExpired:        "token授权已过期",
	CasbinFailed:        "权限不足",
	RegisterFailed:      "注册失败",
	CaptchaVerifyFailed: "验证码校验失败",
	MenuListFailed:      "获取路由菜单失败",
	CasbinAddFailed:     "权限添加失败",
	CasbinDelFailed:     "权限删除失败",
	CasbinUpdateFailed:  "权限更新失败",
	CasbinListFailed:    "权限列表失败",
	RateLimited:         "超出请求频率限制",
	FileWithExcelFailed: "不是excel文件",
	FileReportFailed:    "文件上传失败",
	FileOpenFailed:      "文件打开失败",
	GetUserInfoFailed:   "获取用户信息失败",
	SetUserInfoFailed:   "更新用户信息失败",
	GetCasbinListFailed: "获取权限表信息失败",
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
	c.JSON(http.StatusOK, resp)
}

func Ok(c *gin.Context, data interface{}) {
	if data == nil {
		data = make(map[string]string, 0)
	}
	Result(c, CommonOK, data, "")
}

func Fail(c *gin.Context, code int, err interface{}) {
	data := make(map[string]string, 0)
	Result(c, code, data, err)
}
