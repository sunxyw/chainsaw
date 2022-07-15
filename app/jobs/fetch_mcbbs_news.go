package jobs

import (
	"gohub/app/models/news"

	"github.com/gocolly/colly/v2"
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
		title := e.ChildText("a.s.xst")
		href := e.ChildAttr("a.s.xst", "href")
		href = "https://www.mcbbs.net/" + href

		news.AddNews(title, href)
	})

	c.Visit("https://www.mcbbs.net/forum-news-1.html")
}

func (job *FetchMCBBSNews) ShouldRunAtStartup() bool {
	return true
}

func (job *FetchMCBBSNews) CronSpec() string {
	return "@every 5m"
}
