package v1

import (
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/attribute"

	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/plugins/tracing"
	"github.com/xiaohubai/go-layout/service"
)

func GetDict(c *gin.Context) {
	/* span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()
	span.LogFields(log.Object("service.GetDictList(c)", ""), log.Object("error", err)) */

	ctx, span := tracing.NewSpan(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(ctx)
	defer span.End()

	span.SetAttributes(attribute.Key("ssss").String("hdsgjfd"))

	if dictResp, err := service.GetDictList(c); err != nil {
		response.Fail(c, response.GetDictListFailed, err)
	} else {
		response.Ok(c, dictResp)
	}
}
