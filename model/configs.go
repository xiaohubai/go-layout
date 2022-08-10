package model

type Config struct {
	JWT     JWT                 `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Zap     Zap                 `mapstructure:"zap" json:"zap" yaml:"zap"`
	Redis   Redis               `mapstructure:"redis" json:"redis" yaml:"redis"`
	Email   Email               `mapstructure:"email" json:"email" yaml:"email"`
	Casbin  Casbin              `mapstructure:"casbin" json:"casbin" yaml:"casbin"`
	System  System              `mapstructure:"system" json:"system" yaml:"system"`
	Captcha Captcha             `mapstructure:"captcha" json:"captcha" yaml:"captcha"`
	Mysql   Mysql               `mapstructure:"mysql" json:"mysql" yaml:"mysql"`
	Local   Local               `mapstructure:"local" json:"local" yaml:"local"`
	Excel   Excel               `mapstructure:"excel" json:"excel" yaml:"excel"`
	Kafka   Kafka               `mapstructure:"kafka" json:"kafka" yaml:"kafka"`
	Es      Es                  `mapstructure:"es" json:"es" yaml:"es"`
	Jaeger  Jaeger              `mapstructure:"jaeger" json:"jaeger" yaml:"jaeger"`
	Consul  map[string]Register `mapstructure:"consul" json:"consul" yaml:"consul"`
}
type Register struct {
	Address     string `mapstructure:"address" json:"address" yaml:"address"`
	Scheme      string `mapstructure:"scheme" json:"scheme" yaml:"scheme"`
	HealthCheck bool   `mapstructure:"healthCheck" json:"healthCheck" yaml:"healthCheck"`
	Endpoint    string `mapstructure:"endpoint" json:"endpoint" yaml:"endpoint"`
}

type Jaeger struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`
	Address string `mapstructure:"address" json:"address" yaml:"address"`
}
type Es struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"`
}
type Kafka struct {
	Address []string `mapstructure:"address" json:"address" yaml:"address"` // 地址
	Topics  string   `mapstructure:"topics" json:"topics" yaml:"topics"`    // topics
}

type JWT struct {
	SigningKey  string `mapstructure:"signing-key" json:"signingKey" yaml:"signing-key"`    // jwt签名
	ExpiresTime int64  `mapstructure:"expires-time" json:"expiresTime" yaml:"expires-time"` // 过期时间
	BufferTime  int64  `mapstructure:"buffer-time" json:"bufferTime" yaml:"buffer-time"`    // 缓冲时间
}

type Zap struct {
	Level         string `mapstructure:"level" json:"level" yaml:"level"`                           // 级别
	Format        string `mapstructure:"format" json:"format" yaml:"format"`                        // 输出
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                        // 日志前缀
	Director      string `mapstructure:"director" json:"director"  yaml:"director"`                 // 日志文件夹
	LinkName      string `mapstructure:"link-name" json:"linkName" yaml:"link-name"`                // 软链接名称
	ShowLine      bool   `mapstructure:"show-line" json:"showLine" yaml:"showLine"`                 // 显示行
	EncodeLevel   string `mapstructure:"encode-level" json:"encodeLevel" yaml:"encode-level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace-key" json:"stacktraceKey" yaml:"stacktrace-key"` // 栈名
	LogInConsole  bool   `mapstructure:"log-in-console" json:"logInConsole" yaml:"log-in-console"`  // 输出控制台
}

type Redis struct {
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`                   // redis的哪个数据库
	Addr     string `mapstructure:"addr" json:"addr" yaml:"addr"`             // 服务器地址:端口
	Password string `mapstructure:"password" json:"password" yaml:"password"` // 密码
}

type Email struct {
	To       string `mapstructure:"to" json:"to" yaml:"to"`                   // 收件人:多个以英文逗号分隔
	Port     int    `mapstructure:"port" json:"port" yaml:"port"`             // 端口
	From     string `mapstructure:"from" json:"from" yaml:"from"`             // 收件人
	Host     string `mapstructure:"host" json:"host" yaml:"host"`             // 服务器地址
	IsSSL    bool   `mapstructure:"is-ssl" json:"isSSL" yaml:"is-ssl"`        // 是否SSL
	Secret   string `mapstructure:"secret" json:"secret" yaml:"secret"`       // 密钥
	Nickname string `mapstructure:"nickname" json:"nickname" yaml:"nickname"` // 昵称
}

type Casbin struct {
	ModelPath string `mapstructure:"model-path" json:"modelPath" yaml:"model-path"` // 存放casbin模型的相对路径
}

type System struct {
	Name    string `mapstructure:"name" json:"name" yaml:"name"`             // 服务名
	Version string `mapstructure:"version" json:"version" yaml:"version"`    // 版本
	Port    string `mapstructure:"port" json:"port" yaml:"port"`             // 端口值
	DbType  string `mapstructure:"db-type" json:"dbType" yaml:"db-type"`     // 数据库类型:mysql(默认)|sqlite|sqlserver|postgresql
	GinMode string `mapstructure:"gin-mode" json:"gin-mode" yaml:"gin-mode"` // ginMode模式
	Rate    int    `mapstructure:"rate" json:"rate" yaml:"rate"`             // 限流
}

type Captcha struct {
	KeyLong   int `mapstructure:"key-long" json:"keyLong" yaml:"key-long"`       // 验证码长度
	ImgWidth  int `mapstructure:"img-width" json:"imgWidth" yaml:"img-width"`    // 验证码宽度
	ImgHeight int `mapstructure:"img-height" json:"imgHeight" yaml:"img-height"` // 验证码高度
}

type Mysql struct {
	Path         string `mapstructure:"path" json:"path" yaml:"path"`                             // 服务器地址:端口
	Config       string `mapstructure:"config" json:"config" yaml:"config"`                       // 高级配置
	Dbname       string `mapstructure:"db-name" json:"dbname" yaml:"db-name"`                     // 数据库名
	Username     string `mapstructure:"username" json:"username" yaml:"username"`                 // 数据库用户名
	Password     string `mapstructure:"password" json:"password" yaml:"password"`                 // 数据库密码
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns" yaml:"max-idle-conns"` // 空闲中的最大连接数
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns" yaml:"max-open-conns"` // 打开到数据库的最大连接数
	LogMode      string `mapstructure:"log-mode" json:"logMode" yaml:"log-mode"`                  // 是否开启Gorm全局日志
	LogZap       bool   `mapstructure:"log-zap" json:"logZap" yaml:"log-zap"`                     // 是否通过zap写入日志文件
}

type Local struct {
	Path string `mapstructure:"path" json:"path" yaml:"path"` // 本地文件路径
}

type Excel struct {
	Dir string `mapstructure:"dir" json:"dir" yaml:"dir"`
}
