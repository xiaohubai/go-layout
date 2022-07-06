package v1

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	consulAPI "github.com/hashicorp/consul/api"
	v1 "github.com/xiaohubai/rpc_layout/api/dict/v1"
)

func GetDict(c *gin.Context) {

	cfg := consulAPI.DefaultConfig()
	cfg.Address = "127.0.0.1:8500"
	cfg.Scheme = "http"
	cli, err := consulAPI.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(false))
	endpoint := "discovery://default/rpc_layout"
	conn, err := grpc.DialInsecure(context.Background(), grpc.WithEndpoint(endpoint), grpc.WithDiscovery(r))
	if err != nil {
		panic(err)
	}

	client := v1.NewDictClient(conn)
	resp, _ := client.Get(context.Background(), &v1.DictRequest{Tag: "ssss"})
	fmt.Println(resp)

}
