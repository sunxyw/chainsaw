// Package config 站点配置信息
package config

import "gohub/pkg/config"

func init() {
	config.Add("cache", func() map[string]interface{} {
		return map[string]interface{}{

			// 缓存驱动，支持：memory、redis
			"driver": config.Env("CACHE_DRIVER", "memory"),
		}
	})
}
