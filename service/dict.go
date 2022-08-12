package service

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/xiaohubai/go-layout/configs/consts"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
	"github.com/xiaohubai/go-layout/plugins/consul"
	v1 "github.com/xiaohubai/rpc_layout/api/dict/v1"
)

// GetDictList 获取字典序
func GetDictList(c *gin.Context) (map[string]interface{}, error) {
	r := global.Cfg.Consul[consts.RPCRemoteByRPCLayout]
	d := consul.NewDiscovery(&model.Register{Address: r.Address, Scheme: r.Scheme, HealthCheck: r.HealthCheck})
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint(r.Endpoint),
		grpc.WithDiscovery(d),
		grpc.WithMiddleware(
			tracing.Client(),
		),
	)
	if err != nil {
		return nil, fmt.Errorf("请求服务发现失败")
	}
	proxy := v1.NewDictClient(conn)
	rpcReq := &v1.DictRequest{
		Tag: "ssss",
	}
	rpcResp, err := proxy.Get(c.Request.Context(), rpcReq)
	if err != nil {
		return nil, fmt.Errorf("远程调用失败")
	}
	resp := map[string]interface{}{
		"guide": rpcResp.Guide,
	}
	return resp, nil
}
