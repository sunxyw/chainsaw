package cmd

import (
	"gohub/pkg/logger"
	"gohub/pkg/rcon"

	"github.com/spf13/cobra"
)

var CmdPlay = &cobra.Command{
	Use:   "play",
	Short: "Likes the Go Playground, but running at our application context",
	Run:   runPlay,
}

// 调试完成后请记得清除测试代码
func runPlay(cmd *cobra.Command, args []string) {
	rcon.InitRconClient()
	rcon.AddServer(rcon.ServerConf{
		Name:     "test",
		Host:     "",
		Password: "",
	})
	resp, err := rcon.Server("test").Send("list")
	logger.LogIf(err)
	logger.Dump(resp)
}
