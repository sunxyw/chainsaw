package bungee

import (
	"encoding/json"
	"fmt"
	"gohub/pkg/logger"
	"gohub/pkg/redis"
	"sync"
	"time"
)

type BungeeCluster struct {
	RedisClient *redis.RedisClient
	proxies     []*BungeeProxy
	LastFetch   time.Time
	Lock        sync.RWMutex
}

var Cluster *BungeeCluster

func InitBungeeCluster(redisConf redis.RedisConf) {
	Cluster = &BungeeCluster{
		RedisClient: redis.NewClientWithConf(redisConf),
		proxies:     []*BungeeProxy{},
		LastFetch:   time.Now().Add(-time.Hour),
		Lock:        sync.RWMutex{},
	}
}

func (b *BungeeCluster) FetchProxies() {
	proxyNames := b.RedisClient.HKeys("heartbeats")
	proxies := make([]*BungeeProxy, len(proxyNames))
	for i, name := range proxyNames {
		proxies[i] = NewBungeeProxy(name)
	}
	b.proxies = proxies
}

func (b *BungeeCluster) GetProxies() []*BungeeProxy {
	return b.proxies
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
