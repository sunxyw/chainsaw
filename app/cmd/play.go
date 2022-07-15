package cmd

import (
	"gohub/app/models/news"
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
	// news.AddNews("testing news", "https://example.org/news/1")
	logger.Dump(news.IsExist("url", "https://example.org/news/1"))
}
