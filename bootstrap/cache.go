// Package bootstrap 启动程序功能
package bootstrap

import (
	"fmt"
	"gohub/pkg/cache"
	"gohub/pkg/config"
	"gohub/pkg/logger"
)

// SetupCache 缓存
func SetupCache() {

	var store cache.Store

	switch config.Get[string]("cache.driver") {
	case "redis":
		store = cache.NewRedisStore(
			fmt.Sprintf("%v:%v", config.Get[string]("redis.host"), config.Get[string]("redis.port")),
			config.Get[string]("redis.username"),
			config.Get[string]("redis.password"),
			config.Get[int]("redis.database_cache"),
		)
	case "memory":
		store = cache.NewMemoryStore()
	default:
		logger.Error("不支持的缓存驱动！")
	}

	cache.InitWithCacheStore(store)
}
