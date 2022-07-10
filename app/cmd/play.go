package cmd

import (
	"gohub/pkg/bungee"
	"gohub/pkg/logger"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	bungee.Cluster.FetchProxies()

	for _, proxy := range bungee.Cluster.Proxies {
		logger.InfoString("play", "proxy", proxy.Name)
		proxy.FetchPlayerlist()
		for server, players := range proxy.GetPlayerlist() {
			logger.InfoString("play", "server", server)
			for _, player := range players {
				logger.InfoString("play", "player", player.Name)
			}
		}
	}
}
