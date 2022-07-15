package jobs

import (
	"gohub/app/models/news"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/samber/lo"
)

type FetchMCBBSNews struct {
}

func (job *FetchMCBBSNews) Name() string {
	return "fetch_mcbbs_news"
}

func (job *FetchMCBBSNews) Run() {
	c := colly.NewCollector(colly.AllowedDomains("www.mcbbs.net"))

	// 遍历所有主题
	c.OnHTML(`#threadlisttableid > tbody[id^="normalthread_"]`, func(e *colly.HTMLElement) {
		if lo.Contains(e.ChildAttrs("img", "alt"), "过期") {
			return
		}

		title := e.ChildText("a.s.xst")
		href := e.ChildAttr("a.s.xst", "href")
		href = "https://www.mcbbs.net/" + href

		news.AddNews(title, href, time.Now())
	})

	c.Visit("https://www.mcbbs.net/forum-news-1.html")
}

func (job *FetchMCBBSNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCBBSNews) CronSpec() string {
	return "@every 5m"
}
