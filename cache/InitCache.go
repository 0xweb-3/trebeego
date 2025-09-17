package cache

import (
	"context"
	"fmt"
	"sync"

	"github.com/beego/beego/v2/server/web"
	redis "github.com/redis/go-redis/v9"
)

var (
	client *redis.Client
	once   sync.Once
)

// GetRedisClient 返回全局唯一的 Redis 客户端实例
func GetRedisClient() *redis.Client {
	once.Do(func() {
		// 从 app.conf 读取配置
		addr, _ := web.AppConfig.String("redis.conn")         // e.g. "127.0.0.1:6379"
		password, _ := web.AppConfig.String("redis.password") // e.g. ""
		db, _ := web.AppConfig.Int("redis.dbNum")             // e.g. 0

		// 初始化 Redis 客户端
		client = redis.NewClient(&redis.Options{
			Addr:     addr,
			Password: password,
			DB:       db,
			// go-redis 内部自带连接池，不需要额外写 maxIdle/idleTimeout
		})

		// 测试连接
		if err := client.Ping(context.Background()).Err(); err != nil {
			panic(fmt.Sprintf("Redis connect failed: %v", err))
		}
	})
	return client
}
