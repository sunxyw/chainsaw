package news

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/logger"
	"gohub/pkg/paginator"
	"net/url"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (news News) {
	database.DB.Where("id", idstr).First(&news)
	return
}

func GetBy(field, value string) (news News) {
	database.DB.Where("? = ?", field, value).First(&news)
	return
}

func All() (news []News) {
	database.DB.Find(&news)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(News{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (news []News, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(News{}),
		&news,
		app.V1URL(database.TableName(&News{})),
		perPage,
	)
	return
}

func AddNews(title string, newsURL string) {
	if IsExist("url", newsURL) {
		return
	}

	parsed, err := url.Parse(newsURL)
	logger.LogIf(err)

	news := News{
		Title:  title,
		URL:    newsURL,
		Source: parsed.Hostname(),
	}
	news.Create()
}
