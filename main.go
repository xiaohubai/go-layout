package main

import (
	"context"
	"log"
	"time"

	"github.com/xiaohubai/go-layout/configs/global"
	_ "github.com/xiaohubai/go-layout/plugins"
	"github.com/xiaohubai/go-layout/plugins/tracing"
	"github.com/xiaohubai/go-layout/router"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
)

func main() {
	tp := tracing.OpentelemetryInit()
	defer tp.Shutdown(context.Background())

	if err := HTTPServer(); err != nil {
		log.Fatal(err)
	}

}

//HTTPServer 服务启动
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
