package main

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"github.com/xiaohubai/go-layout/configs/global"
	_ "github.com/xiaohubai/go-layout/plugins"
	"github.com/xiaohubai/go-layout/plugins/ants"
	"github.com/xiaohubai/go-layout/plugins/tracing"
	"github.com/xiaohubai/go-layout/router"
)

func main() {
	closer := tracing.Init()
	defer closer.Close()
	defer ants.Release()
	wg := sync.WaitGroup{}
	defer wg.Wait()

	wg.Add(1)
	ants.Go(func() {
		if err := HTTPServer(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	})

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
