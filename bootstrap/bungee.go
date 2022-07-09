package bootstrap

import (
	"gohub/pkg/bungee"
	"gohub/pkg/config"
	"gohub/pkg/redis"
)

func SetupBungee() {
	bungee.InitBungeeCluster(config.Get[redis.RedisConf]("redis.bungee"))
}
