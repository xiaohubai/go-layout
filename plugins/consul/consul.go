package consul

import (
	"fmt"
	"time"

	"github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/registry"
	consulAPI "github.com/hashicorp/consul/api"
	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/model"
)

func Config(remoteName, congifType string, result interface{}) (*viper.Viper, error) {
	v := viper.New()
	v.AddRemoteProvider(global.ConsulName, global.ConsulURL, remoteName)
	v.SetConfigType(congifType)
	if err := v.ReadRemoteConfig(); err != nil {
		return nil, err
	}
	if err := v.Unmarshal(result); err != nil {
		return nil, err
	}
	go func() {
		for {
			time.Sleep(time.Second * 5)
			err := v.WatchRemoteConfig()
			if err != nil {
				fmt.Printf("unable to read remote config: %v", err)
				continue
			}
			v.Unmarshal(result)
		}
	}()
	return v, nil
}

func NewDiscovery(conf *model.Register) registry.Discovery {
	c := consulAPI.DefaultConfig()
	c.Address = conf.Address
	c.Scheme = conf.Scheme
	cli, err := consulAPI.NewClient(c)
	if err != nil {
		panic(err)
	}
	r := consul.New(cli, consul.WithHealthCheck(conf.HealthCheck))
	return r
}
