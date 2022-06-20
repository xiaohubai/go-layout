# 代码结构
``` go
.
├── Dockerfile
├── Makefile
├── README.MD
├── api
│   └── v1
│       ├── captcha.go
│       ├── token.go
│       └── user.go
├── config.docker.yaml
├── config_dev.yaml
├── configs
│   ├── consts
│   │   └── consts.go
│   └── global
│       └── global.go
├── dao
│   ├── dao.go
│   └── user.go
├── deploy
│   ├── docker-compose.yml
│   └── volumes
│       ├── README..md
│       ├── elasticsearch
│       │   ├── data
│       │   └── elasticsearch.yml
│       ├── grafana
│       │   ├── data
│       │   ├── datasources
│       │   │   └── prometheus_ds.yml
│       │   └── grafana.ini
│       ├── kibana
│       │   └── kibana.yml
│       ├── mysql
│       │   └── data
│       ├── prometheus
│       │   └── prometheus.yml
│       └── redis
│           ├── data
│           └── redis.conf
├── docs
├── go.mod
├── go.sum
├── main.go
├── middleware
│   ├── casbin.go
│   ├── cors.go
│   ├── jaeger.go
│   ├── jwt.go
│   ├── metrics.go
│   ├── redis_limiter.go
│   ├── request.go
│   └── translations.go
├── model
│   ├── configs.go
│   ├── jwt.go
│   ├── menu.go
│   ├── request
│   │   ├── token.go
│   │   └── user.go
│   ├── response
│   │   ├── captcha.go
│   │   ├── response.go
│   │   └── user.go
│   └── user.go
├── plugins
│   ├── ants
│   │   └── ants.go
│   ├── es
│   │   └── es.go
│   ├── gorm
│   │   ├── gorm.go
│   │   └── logger.go
│   ├── jwt
│   │   └── jwt.go
│   ├── metrics
│   │   └── metrics.go
│   ├── plugins.go
│   ├── redis
│   │   └── redis.go
│   ├── tables
│   │   └── tables.go
│   ├── tracing
│   │   └── tracing.go
│   ├── viper
│   │   └── viper.go
│   └── zap
│       └── zap.go
├── rbac_model.conf
├── router
│   └── router.go
├── scripts
│   └── start.sh
├── service
│   ├── token.go
│   └── user.go
└── utils
    ├── math.go
    ├── path.go
    ├── request.go
    ├── time.go
    ├── token.go
    ├── trace_id.go
    └── validator.go

```

## 开发过程


## TODO
1.发现painc，提取错误堆栈，存放到日志，优先中间件


