package jobs

type Job interface {
	Name() string
	Run()
	ShouldRunAtStartup() bool
	CronSpec() string
}
