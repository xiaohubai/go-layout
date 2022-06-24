package service

import (
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/dao"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/model/request"
	"github.com/xuri/excelize/v2"
)

func AddCasbin(c *gin.Context, r request.CasbinReq) error {
	t := model.CasbinRule{
		Ptype: "p",
		V0:    r.RoleId,
		V1:    r.Path,
		V2:    strings.ToUpper(r.Method),
	}
	err := dao.AddCasbin(c, []model.CasbinRule{t})
	if err != nil {
		return err
	}
	return nil
}

func AddCasbinWithExcel(c *gin.Context, file io.Reader) error {
	f, err := excelize.OpenReader(file)
	if err != nil {
		return err
	}
	defer f.Close()

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	casbins := []model.CasbinRule{}
	for _, row := range rows {
		if row[0] != "" && row[1] != "" && row[2] != "" {
			data := model.CasbinRule{
				//	Ptype: "p",
				V0: row[0],
				V1: row[1],
				V2: strings.ToUpper(row[2]),
			}
			casbins = append(casbins, data)
		}
	}
	err = dao.AddCasbin(c, casbins)
	if err != nil {
		return err
	}
	return nil
}
