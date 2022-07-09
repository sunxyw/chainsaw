package bungee

import "gohub/pkg/redis"

type BungeeCluster struct {
	RedisClient *redis.RedisClient
	Proxies     []BungeeProxy
}

var Cluster *BungeeCluster

func InitBungeeCluster(redisConf redis.RedisConf) {
	Cluster = &BungeeCluster{
		RedisClient: redis.NewClientWithConf(redisConf),
	}
}

func (b *BungeeCluster) FetchProxies() {
	proxyNames := b.RedisClient.HKeys("heartbeats")
	proxies := make([]BungeeProxy, len(proxyNames))
	for _, name := range proxyNames {
		proxies = append(proxies, BungeeProxy{Name: name})
	}
	b.Proxies = proxies
}
