package viper

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/plugins/consul"
)

// Init 读取/实时检测配置组件
func Init() *viper.Viper {
	var conf string
	flag.StringVar(&conf, "c", "", "choose config")
	flag.Parse()
	if conf != "" {
		v, err := consul.Config("config/"+conf, "yaml", &global.Cfg)
		if err != nil {
			panic(err)
		}
		return v
	} else {
		conf = global.ConfigFileEnv
		v := viper.New()
		v.SetConfigFile(conf)
		v.SetConfigType("yaml")
		if err := v.ReadInConfig(); err != nil {
			panic(err)
		}
		v.WatchConfig()
		v.OnConfigChange(func(e fsnotify.Event) {
			if err := v.Unmarshal(&global.Cfg); err != nil {
				fmt.Println("err:", err)
			}
		})
		if err := v.Unmarshal(&global.Cfg); err != nil {
			fmt.Println("err:", err)
		}
		return v
	}

}
