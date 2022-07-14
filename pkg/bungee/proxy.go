package bungee

import (
	"fmt"
	"gohub/pkg/logger"
)

type BungeeProxy struct {
	Name       string
	playerlist map[string][]BungeePlayer
}

func NewBungeeProxy(name string) *BungeeProxy {
	return &BungeeProxy{
		Name:       name,
		playerlist: make(map[string][]BungeePlayer),
	}
}

func (b *BungeeProxy) GetPlayerlist() map[string][]BungeePlayer {
	waitUntilFetchFinished()
	return b.playerlist
}

func (b *BungeeProxy) FetchPlayerlist() {
	uuids := Cluster.RedisClient.SMembers(fmt.Sprintf("proxy:%v:usersOnline", b.Name))
	if len(uuids) == 0 {
		return
	}

	names := Cluster.GetCachedPlayerNames(uuids)

	list := make(map[string][]BungeePlayer)
	for _, uuid := range uuids {
		playerInfo := Cluster.GetPlayerInfo(uuid)

		if playerInfo["online"] != "0" {
			logger.WarnString("bungee", "playerlist", "player not online while sync: "+uuid)
			continue
		}

		player := BungeePlayer{
			UUID:   uuid,
			Name:   names[uuid],
			IP:     playerInfo["ip"],
			Server: playerInfo["server"],
		}

		list[player.Server] = append(list[player.Server], player)
	}

	b.playerlist = list
}
