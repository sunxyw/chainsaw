//Package notification 模型
package notification

import (
	"gohub/app/models"
	"gohub/pkg/database"
	"time"
)

type Notification struct {
	models.BaseModel

	Type   string    `json:"type"`
	Data   string    `json:"data"`
	ReadAt time.Time `json:"read_at"`

	models.CommonTimestampsField
}

func (notification *Notification) Create() {
	database.DB.Create(&notification)
}

func (notification *Notification) Save() (rowsAffected int64) {
	result := database.DB.Save(&notification)
	return result.RowsAffected
}

func (notification *Notification) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&notification)
	return result.RowsAffected
}

func (notification *Notification) MarkAsRead() (rowsAffected int64) {
	notification.ReadAt = time.Now()
	return notification.Save()
}
