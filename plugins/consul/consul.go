package consul

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
	_ "github.com/spf13/viper/remote"
)

func Config(remoteName, congifType string, result interface{}) (*viper.Viper, error) {
	v := viper.New()
	v.AddRemoteProvider("consul", "172.21.0.2:8500", remoteName)
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
