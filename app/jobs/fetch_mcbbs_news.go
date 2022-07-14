package jobs

import (
	"github.com/gocolly/colly/v2"
)

type FetchMCBBSNews struct {
	Result []thread
}

type thread struct {
	Title string
	URL   string
}

func (job *FetchMCBBSNews) Name() string {
	return "fetch_mcbbs_news"
}

func (job *FetchMCBBSNews) Run() {
	c := colly.NewCollector()

	job.Result = make([]thread, 0)

	// 遍历所有主题
	c.OnHTML(`#threadlisttableid > tbody[id^="normalthread_"]`, func(e *colly.HTMLElement) {
		if len(job.Result) >= 5 {
			return
		}
		title := e.ChildText("a.s.xst")
		href := e.ChildAttr("a.s.xst", "href")
		job.Result = append(job.Result, thread{Title: title, URL: href})
	})

	c.Visit("https://www.mcbbs.net/forum-news-1.html")
}

func (job *FetchMCBBSNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCBBSNews) CronSpec() string {
	return "@every 5m"
}
