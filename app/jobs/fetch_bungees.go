package jobs

import (
	"gohub/pkg/bungee"
	"time"
)

type FetchBungees struct {
}

func (job *FetchBungees) Name() string {
	return "fetch_bungees"
}

func (job *FetchBungees) Run() {
	bungee.Cluster.Lock.Lock()
	defer bungee.Cluster.Lock.Unlock()

	bungee.Cluster.FetchProxies()

	for _, proxy := range bungee.Cluster.GetProxies() {
		proxy.FetchPlayerlist()
	}

	bungee.Cluster.LastFetch = time.Now()
}

func (job *FetchBungees) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchBungees) CronSpec() string {
	return "@every 30s"
}
