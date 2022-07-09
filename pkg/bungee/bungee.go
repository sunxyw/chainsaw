package bungee

import "gohub/pkg/redis"

type BungeeRedis struct {
	RedisClient *redis.RedisClient
	Proxies     []BungeeProxy
}

func NewBungeeRedis(redisClient *redis.RedisClient) *BungeeRedis {
	return &BungeeRedis{
		RedisClient: redisClient,
	}
}

func (b *BungeeRedis) FetchProxies() {
	proxyNames := b.RedisClient.HKeys("heartbeats")
	proxies := make([]BungeeProxy, len(proxyNames))
	for _, name := range proxyNames {
		proxies = append(proxies, BungeeProxy{Name: name})
	}
	b.Proxies = proxies
}
