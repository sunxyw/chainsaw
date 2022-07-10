package jobs

import (
	"gohub/pkg/bungee"
	"gohub/pkg/logger"
	"time"
)

type FetchBungees struct {
}

func (job *FetchBungees) Run() {
	bungee.Cluster.FetchProxies()

	for _, proxy := range bungee.Cluster.Proxies {
		proxy.FetchPlayerlist()
	}

	bungee.Cluster.LastFetch = time.Now()

	logger.InfoString("cronjob", "bungee", "playerlist fetched")
}
