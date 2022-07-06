# 代码结构
``` go
.
├── deploy
│   ├── docker-compose.yml
│   └── volumes
├── configs
├── docs
├── api
├── dao
├── model
├── service
├── middleware
├── plugins
├── router
├── scripts
├── utils
├── rbac_model.conf
├── config.docker.yaml
├── config_dev.yaml
├── Dockerfile
├── Makefile
├── main.go
├── go.mod
├── go.sum
└── README.MD

```

## 主要功能
- gin web框架 （github.com/gin-gonic/gin）
- jwt认证（github.com/golang-jwt/jwt）
- casbin鉴权（github.com/casbin/casbin/v2）
- gorm数据库组件（gorm.io/gorm）
- viper实时解析检测配置文件（github.com/spf13/viper）
- swagger 接口文档生成 （github.com/swaggo/swag）
- redis组件 （github.com/go-redis/redis）
- zap日志定制化 （go.uber.org/zap）
- 参数校验（github.com/go-playground/validator/v10）
- 参数校验错误信息翻译（github.com/go-playground/universal-translator）
- kafka组件 （github.com/Shopify/sarama）
- es组件 （github.com/olivere/elastic/v7）
- email 组件 （github.com/jordan-wright/email）
- jaeger全链路 （github.com/uber/jaeger-client-go）
- prometheus监控 （github.com/prometheus/client_golang）
- excel处理 （github.com/xuri/excelize/v2）
- 协程控制 （github.com/panjf2000/ants/v2）
- 分布式接口限流 （github.com/go-redis/redis_rate/v9）

## 一键搭建环境
- 拷贝volumes文件到/usr/local/
- 增加内网ip：172.12.0.2
- 执行 make compose

#TODO
1.新增rpc服务，获取kafka消息，topic告警发邮件
