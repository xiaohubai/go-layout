package v1

import (
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
	"github.com/xiaohubai/go-layout/utils"
)

func AddCasbin(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()

	var r request.CasbinReq
	if err := utils.ShouldBindJSON(c, &r); err != nil {
		span.LogFields(log.Object("utils.ShouldBindJSON(c, &r)", r), log.Object("error", err))
		response.Fail(c, response.ParamsFail, err)
		return
	}

	if err := service.AddCasbin(c, r); err != nil {
		span.LogFields(log.Object("service.AddCasbin(c, r)", r), log.Object("error", err))
		response.Fail(c, response.CasbinAddFail, nil)
	} else {
		response.Ok(c, nil)
	}

}

func AddCasbinWithExcel(c *gin.Context) {
	span, ctx := opentracing.StartSpanFromContext(c.Request.Context(), "api")
	c.Request = c.Request.WithContext(opentracing.ContextWithSpan(ctx, span))
	defer span.Finish()
	file, fh, err := c.Request.FormFile("file")
	if err != nil {
		response.Fail(c, response.ErrFileReport, nil)
		return
	}
	defer file.Close()

	if !strings.HasSuffix(fh.Filename, "xlsx") {
		response.Fail(c, response.ErrFileWithExcel, nil)
		return
	}
	if err := service.AddCasbinWithExcel(c, file); err != nil {
		span.LogFields(log.Object("service.AddCasbinWithExcel(c, file)", ""), log.Object("error", err))
		response.Fail(c, response.CasbinAddFail, nil)
	} else {
		response.Ok(c, nil)
	}

}
