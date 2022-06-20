package viper

import (
	"flag"
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-layout/configs/global"
)

// Init 读取/实时检测配置组件
func Init(path ...string) *viper.Viper {
	var conf string
	if len(path) == 0 {
		//默认命令行 -c参数传递
		flag.StringVar(&conf, "c", "", "choose config file.")
		flag.Parse()
		if conf == "" {
			conf = global.ConfigEnv
		}
	} else {
		//使用func Viper()传递的值
		conf = path[0]
	}

	v := viper.New()
	v.SetConfigFile(conf)
	v.SetConfigType("yaml")
	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("viper register failed: %s \n", err))
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
