package global

import (
	"github.com/Shopify/sarama"
	"github.com/go-redis/redis/v8"
	"github.com/hpcloud/tail"
	"github.com/olivere/elastic/v7"
	"github.com/opentracing/opentracing-go"
	"github.com/spf13/viper"
	"github.com/xiaohubai/go-layout/model"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	ConsulName    = "consul"
	ConsulURL     = "172.21.0.2:8500"
	ConfigFileEnv = "./config_dev.yaml"
)
var (
	Viper  *viper.Viper
	Db     *gorm.DB
	Redis  *redis.Client
	Es     *elastic.Client
	Log    *zap.Logger
	Cfg    model.Config
	Tracer opentracing.Tracer
)

var (
	KafkaProducer sarama.SyncProducer
	KafkaConsumer sarama.Consumer
)

var (
	TailObj *tail.Tail
	LogChan chan string
)
