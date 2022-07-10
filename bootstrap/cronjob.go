package bootstrap

import (
	"gohub/app/jobs"
	"gohub/pkg/logger"

	cronpkg "github.com/robfig/cron/v3"
)

func SetupCronjob() {
	cron := cronpkg.New()
	cron.AddJob("@every 1m", &jobs.FetchBungees{})
	cron.Start()
	logger.InfoString("cronjob", "scheduler", "started")
}
