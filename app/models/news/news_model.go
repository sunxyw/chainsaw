//Package new 模型
package news

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type News struct {
	models.BaseModel

	Title  string `json:"title"`
	URL    string `json:"url"`
	Source string `json:"source"`

	models.CommonTimestampsField
}

func (news *News) Create() {
	database.DB.Create(&news)
}

func (news *News) Save() (rowsAffected int64) {
	result := database.DB.Save(&news)
	return result.RowsAffected
}

func (news *News) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&news)
	return result.RowsAffected
}
