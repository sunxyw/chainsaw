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
	proxies     []*BungeeProxy
	LastFetch   time.Time
	Fetching    bool
}

var Cluster *BungeeCluster

func InitBungeeCluster(redisConf redis.RedisConf) {
	Cluster = &BungeeCluster{
		RedisClient: redis.NewClientWithConf(redisConf),
		proxies:     []*BungeeProxy{},
		LastFetch:   time.Now().Add(-time.Hour),
		Fetching:    true,
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

func (b *BungeeCluster) GetProxies(force ...bool) []*BungeeProxy {
	if len(force) == 0 || !force[0] {
		waitUntilFetchFinished()
	}
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

func waitUntilFetchFinished() {
	for Cluster.Fetching {
		time.Sleep(100 * time.Millisecond)
	}
}
