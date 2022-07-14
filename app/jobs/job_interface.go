package jobs

type Job interface {
	Run()
	ShouldRunAtStartup() bool
	CronSpec() string
}
