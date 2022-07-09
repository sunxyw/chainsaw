package cmd

import (
	"gohub/pkg/hash"
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
	password := "testing_passwd"
	hashed := "$2a$14$vdhjHk5aAWJdt0MAOMXPGeB9LvqwbgDSLAubggUSAlJflSk/CXqnS"
	logger.Dump(password)
	logger.Dump(hashed)
	logger.Dump(hash.BcryptCheck(password, hashed))
}
