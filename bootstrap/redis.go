package bootstrap

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {

	if !config.Get[bool]("redis.enable") {
		return
	}

	mainConf := config.Get[redis.RedisConf]("redis.default")
	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", mainConf.Host, mainConf.Port),
		"",
		mainConf.Password,
		mainConf.Database,
	)
}
