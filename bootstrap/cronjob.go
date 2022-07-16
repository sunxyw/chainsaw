package bootstrap

import (
	"gohub/app/jobs"
	"gohub/pkg/logger"
	"time"

	cronpkg "github.com/robfig/cron/v3"
	"go.uber.org/zap"
)

var cron *cronpkg.Cron
var registerdJobs []jobs.Job

func SetupCronjob() {
	cron = cronpkg.New()

	registerJob(&jobs.FetchBungees{})
	registerJob(&jobs.FetchMCBBSNews{})
	registerJob(&jobs.FetchMCVersionsAsNews{})

	cron.Start()
	logger.Info("Cronjob Scheduler Started", zap.Int("jobs", len(registerdJobs)))
}

func registerJob(job jobs.Job) {
	jobFunc := func() {
		start := time.Now()

		job.Run()

		logger.Info("Cronjob Executed", zap.String("job", job.Name()), zap.String("duration", time.Since(start).String()))
	}

	cron.AddFunc(job.CronSpec(), jobFunc)

	if job.ShouldRunAtStartup() {
		jobFunc()
	}

	registerdJobs = append(registerdJobs, job)
}
