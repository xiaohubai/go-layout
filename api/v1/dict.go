package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"

	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
)

func GetDict(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	if dictResp, err := service.GetDictList(c); err != nil {
		span.LogFields(log.Object("service.GetDictList(c)", ""), log.Object("error", err))
		response.Fail(c, response.GetDictListFailed, err)
	} else {
		response.Ok(c, dictResp)
	}
}
