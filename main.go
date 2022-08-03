package main

import (
	"fmt"
	"log"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/configs/global"
	_ "github.com/xiaohubai/go-layout/plugins"
	"github.com/xiaohubai/go-layout/plugins/tracing"
	"github.com/xiaohubai/go-layout/router"
)

func main() {
	t := tracing.Init()
	defer t.Close()

	if err := HTTPServer(); err != nil {
		log.Fatal(err)
	}
}

//Server 服务启动
func HTTPServer() error {
	gin.SetMode(global.Cfg.System.GinMode)
	r := router.Routers()

	port := fmt.Sprintf(":%d", global.Cfg.System.Port)
	s := endless.NewServer(port, r)
	s.ReadHeaderTimeout = 10 * time.Millisecond
	s.WriteTimeout = 10 * time.Second
	s.MaxHeaderBytes = 1 << 20

	err := s.ListenAndServe()
	return err
}
