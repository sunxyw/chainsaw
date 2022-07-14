package jobs

import (
	"gohub/pkg/logger"

	"github.com/gocolly/colly/v2"
)

type FetchMCBBSNews struct {
}

type thread struct {
	Title string
	URL   string
}

func (job *FetchMCBBSNews) Run() {
	c := colly.NewCollector()

	news := make([]thread, 0)

	// 遍历所有主题
	c.OnHTML(`#threadlisttableid > tbody[id^="normalthread_"]`, func(e *colly.HTMLElement) {
		if len(news) >= 5 {
			return
		}
		title := e.ChildText("a.s.xst")
		href := e.ChildAttr("a.s.xst", "href")
		news = append(news, thread{Title: title, URL: href})
	})

	c.Visit("https://www.mcbbs.net/forum-news-1.html")

	logger.InfoString("cronjob", "mcbbs", "news fetched")
}

func (job *FetchMCBBSNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCBBSNews) CronSpec() string {
	return "@every 5m"
}
