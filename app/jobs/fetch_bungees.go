package jobs

import (
	"gohub/pkg/bungee"
	"gohub/pkg/logger"
	"time"
)

type FetchBungees struct {
}

func (job *FetchBungees) Run() {
	bungee.Cluster.Fetching = true

	bungee.Cluster.FetchProxies()

	for _, proxy := range bungee.Cluster.GetProxies(true) {
		proxy.FetchPlayerlist()
	}

	bungee.Cluster.LastFetch = time.Now()
	bungee.Cluster.Fetching = false

	logger.InfoString("cronjob", "bungee", "playerlist fetched")
}

func (job *FetchBungees) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchBungees) CronSpec() string {
	return "@every 20s"
}
