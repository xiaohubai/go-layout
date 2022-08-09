package service

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/plugins/consul"
	v1 "github.com/xiaohubai/rpc_layout/api/dict/v1"
)

func GetDictList(c *gin.Context) (map[string]interface{}, error) {
	r := global.Cfg.Consul[consts.RPCRemoteByRPCLayout]
	d := consul.NewDiscovery(&model.Register{Address: r.Address, Scheme: r.Scheme, HealthCheck: r.HealthCheck})
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(r.Endpoint),
		grpc.WithDiscovery(d),
	)
	if err != nil {
		return nil, fmt.Errorf("远程调用失败")
	}
	client := v1.NewDictClient(conn)
	resp, err := client.Get(context.Background(), &v1.DictRequest{Tag: "ssss"})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp)
	return nil, nil
}
