// Package bootstrap 启动程序功能
package bootstrap

import (
	"gohub/pkg/cache"
	"gohub/pkg/config"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
)

// SetupCache 缓存
func SetupCache() {

	var store cache.Store

	switch config.Get[string]("cache.driver") {
	case "redis":
		store = cache.NewRedisStore(config.Get[redis.RedisConf]("redis.cache"))
	case "memory":
		store = cache.NewMemoryStore()
	default:
		logger.Error("不支持的缓存驱动！")
	}

	cache.InitWithCacheStore(store)
}
