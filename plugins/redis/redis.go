package redis

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/xiaohubai/go-layout/configs/global"
	"github.com/xiaohubai/go-layout/utils"
)

// Redis 内存数据库组件
func Init() *redis.Client {
	redisCfg := global.Cfg.Redis
	client := redis.NewClient(&redis.Options{
		Addr:     redisCfg.Addr,
		Password: redisCfg.Password,
		DB:       redisCfg.DB,
	})
	_, err := client.Ping(client.Context()).Result()
	if err != nil {
		panic(fmt.Errorf("redis connect ping failed: %s \n", err))
	}
	return client
}

// Get 获取字符串
func Get(ctx context.Context, key string) (string, error) {
	return global.Redis.Get(ctx, utils.Md5([]byte(key))).Result()

}

// Set 设置key value 和过期时间   eg：10 * time.Second
func Set(ctx context.Context, key, value string, expiration time.Duration) error {
	return global.Redis.Set(ctx, utils.Md5([]byte(key)), value, expiration).Err()
}

// Del 删除字符串
func Del(ctx context.Context, key string) error {
	return global.Redis.Del(ctx, utils.Md5([]byte(key))).Err()
}
