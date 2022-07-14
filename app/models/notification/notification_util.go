package notification

import (
	"gohub/pkg/app"
	"gohub/pkg/database"
	"gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (notification Notification) {
	database.DB.Where("id", idstr).First(&notification)
	return
}

func GetBy(field, value string) (notification Notification) {
	database.DB.Where("? = ?", field, value).First(&notification)
	return
}

func All() (notifications []Notification) {
	database.DB.Find(&notifications)
	return
}

func IsExist(field, value string) bool {
	var count int64
	database.DB.Model(Notification{}).Where("? = ?", field, value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (notifications []Notification, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(Notification{}),
		&notifications,
		app.V1URL(database.TableName(&Notification{})),
		perPage,
	)
	return
}

func GetOldestUnread() (notification Notification) {
	database.DB.Where("read_at", nil).Order("created_at ASC").First(&notification)
	return
}
