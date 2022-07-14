package bootstrap

import (
	"gohub/app/jobs"
	"gohub/pkg/logger"

	cronpkg "github.com/robfig/cron/v3"
)

var cron *cronpkg.Cron

func SetupCronjob() {
	cron = cronpkg.New()

	registerJob(&jobs.FetchBungees{})
	registerJob(&jobs.FetchMCBBSNews{})

	cron.Start()
	logger.InfoString("cronjob", "scheduler", "started")
}

func registerJob(job jobs.Job) {
	cron.AddJob(job.CronSpec(), job)

	if job.ShouldRunAtStartup() {
		job.Run()
	}
}
