package bootstrap

import (
	"gohub/app/jobs"
	"gohub/pkg/logger"
	"time"

	cronpkg "github.com/robfig/cron/v3"
)

var cron *cronpkg.Cron

func SetupCronjob() {
	cron = cronpkg.New()

	registerJob(&jobs.FetchBungees{})
	registerJob(&jobs.FetchMCBBSNews{})
	registerJob(&jobs.FetchMCVersionsAsNews{})

	cron.Start()
	logger.InfoString("cronjob", "scheduler", "started")
}

func registerJob(job jobs.Job) {
	jobFunc := func() {
		start := time.Now()

		job.Run()

		logger.InfoString("cronjob", job.Name(), "finished in "+time.Since(start).String())
	}

	cron.AddFunc(job.CronSpec(), jobFunc)

	if job.ShouldRunAtStartup() {
		jobFunc()
	}
}
