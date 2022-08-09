package v1

import (
	"github.com/gin-gonic/gin"

	"github.com/xiaohubai/go-layout/model/response"
	"github.com/xiaohubai/go-layout/service"
)

func GetDict(c *gin.Context) {
	if dictResp, err := service.GetDictList(c); err != nil {
		response.Fail(c, response.GetDictListFailed, nil)
	} else {
		response.Ok(c, dictResp)
	}
}
