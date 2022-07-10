package bungee

import (
	"encoding/json"
	"fmt"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"time"
)

type BungeeCluster struct {
	RedisClient *redis.RedisClient
	Proxies     []BungeeProxy
	LastFetch   time.Time
}

var Cluster *BungeeCluster

func InitBungeeCluster(redisConf redis.RedisConf) {
	Cluster = &BungeeCluster{
		RedisClient: redis.NewClientWithConf(redisConf),
		Proxies:     []BungeeProxy{},
		LastFetch:   time.Now().Add(-time.Hour),
	}
}

func (b *BungeeCluster) FetchProxies() {
	proxyNames := b.RedisClient.HKeys("heartbeats")
	proxies := make([]BungeeProxy, len(proxyNames))
	for i, name := range proxyNames {
		proxies[i] = *NewBungeeProxy(name)
	}
	b.Proxies = proxies
}

func (b *BungeeCluster) GetPlayerInfo(uuid string) map[string]string {
	return b.RedisClient.HGetAll(fmt.Sprintf("player:%v", uuid))
}

func (b *BungeeCluster) GetCachedPlayerNames(uuids []string) map[string]string {
	cache := b.RedisClient.HMGet("uuid-cache", uuids...)
	names := make(map[string]string)

	for _, jsonString := range cache {
		var decoded map[string]interface{}
		err := json.Unmarshal([]byte(jsonString), &decoded)
		logger.LogIf(err)

		names[decoded["uuid"].(string)] = decoded["name"].(string)
	}

	return names
}
