package config

import (
	"gohub/pkg/config"
	"gohub/pkg/redis"

	"github.com/spf13/cast"
)

func init() {

	config.Add("redis", func() map[string]interface{} {
		return map[string]interface{}{

			"enable": cast.ToBool(config.Env("REDIS_ENABLE", false)),

			"default": redis.RedisConf{
				Enable:   cast.ToBool(config.Env("REDIS_ENABLE", false)),
				Host:     cast.ToString(config.Env("REDIS_HOST", "127.0.0.1")),
				Port:     cast.ToString(config.Env("REDIS_PORT", "6379")),
				Password: cast.ToString(config.Env("REDIS_PASSWORD", "")),
				Database: cast.ToInt(config.Env("REDIS_MAIN_DB", 1)),
			},

			"cache": redis.RedisConf{
				Enable:   cast.ToBool(config.Env("REDIS_ENABLE", false)),
				Host:     cast.ToString(config.Env("REDIS_HOST", "127.0.0.1")),
				Port:     cast.ToString(config.Env("REDIS_PORT", "6379")),
				Password: cast.ToString(config.Env("REDIS_PASSWORD", "")),
				Database: cast.ToInt(config.Env("REDIS_CACHE_DB", 0)),
			},

			"bungee": redis.RedisConf{
				Enable:   cast.ToBool(config.Env("REDIS_BUNGEE_ENABLE", false)),
				Host:     cast.ToString(config.Env("REDIS_BUNGEE_HOST", "127.0.0.1")),
				Port:     cast.ToString(config.Env("REDIS_BUNGEE_PORT", "6379")),
				Password: cast.ToString(config.Env("REDIS_BUNGEE_PASSWORD", "")),
				Database: cast.ToInt(config.Env("REDIS_BUNGEE_DB", 0)),
			},
		}
	})
}
