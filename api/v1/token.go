package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/metrics"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

func Token(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	var req request.TokenReq
	if err := utils.ShouldBindJSON(c, &req); err != nil {
		span.LogFields(log.Object("utils.ShouldBindJSON(c, &req)", req), log.Object("error", err))
		response.Fail(c, response.ParamsFail, err)
		return
	}
	u := &model.User{Username: req.Username, Password: req.Password}
	resp, err := service.Token(c, u)
	if err != nil {
		span.LogFields(log.Object("service.Token(c, u)", u), log.Object("error", err))
		response.Fail(c, response.TokenFail, err)
		return
	}
	metrics.CounterIncM("token")
	response.Ok(c, resp)
}
