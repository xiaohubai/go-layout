package plugins

import (
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/plugins/es"
	"github.com/xiaohubai/go-layout/plugins/gorm"
	"github.com/xiaohubai/go-layout/plugins/kafka"
	"github.com/xiaohubai/go-layout/plugins/metrics"
	"github.com/xiaohubai/go-layout/plugins/redis"
	"github.com/xiaohubai/go-layout/plugins/tables"
	"github.com/xiaohubai/go-layout/plugins/viper"
	"github.com/xiaohubai/go-layout/plugins/zap"
)

func init() {
	global.Viper = viper.Init()             // 加载环境配置组件
	global.Log = zap.Init()                 // 加载日志组件
	global.Redis = redis.Init()             // 加载redis组件
	global.Db = gorm.Init()                 // 加载数据库组件
	global.Es = es.Init()                   // 加载elasticsearch组件
	global.KafkaProducer = kafka.Producer() // 加载kafka producer
	global.KafkaConsumer = kafka.Consumer() // 加载kafka consumer
	tables.Init()                           // 加载表结构
	metrics.Init()                          // 加载监控指标
}
