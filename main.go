package main

import (
	"log"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/configs/global"
	_ "github.com/xiaohubai/go-layout/plugins"
	"github.com/xiaohubai/go-layout/plugins/jaeger"
	"github.com/xiaohubai/go-layout/router"
)

func main() {
	j := jaeger.Init()
	defer j.Close()
	if err := HTTPServer(); err != nil {
		log.Fatal(err)
	}

}

//Server 服务启动
func HTTPServer() error {
	gin.SetMode(global.Cfg.System.GinMode)
	r := router.Routers()
	s := endless.NewServer(global.Cfg.System.Port, r)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	err := s.ListenAndServe()
	return err
}
