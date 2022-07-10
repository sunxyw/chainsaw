package bungee

import (
	"fmt"
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
	return b.playerlist
}

func (b *BungeeProxy) FetchPlayerlist() {
	uuids := Cluster.RedisClient.SMembers(fmt.Sprintf("proxy:%v:usersOnline", b.Name))
	if len(uuids) == 0 {
		return
	}

	names := Cluster.GetCachedPlayerNames(uuids)
	for _, uuid := range uuids {
		playerInfo := Cluster.GetPlayerInfo(uuid)

		if playerInfo["online"] != "0" {
			continue
		}

		player := BungeePlayer{
			UUID:   uuid,
			Name:   names[uuid],
			IP:     playerInfo["ip"],
			Server: playerInfo["server"],
		}

		b.playerlist[player.Server] = append(b.playerlist[player.Server], player)
	}
}
