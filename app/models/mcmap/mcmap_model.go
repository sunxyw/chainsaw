//Package mcmap 模型
package mcmap

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type Mcmap struct {
	models.BaseModel

	Name    string `json:"name"`
	Type    string `json:"type"`
	Address string `json:"address"`
	Size    int64  `json:"size"`

	models.CommonTimestampsField
}

func (mcmap *Mcmap) TableName() string {
	return "maps"
}

func (mcmap *Mcmap) Create() {
	database.DB.Create(&mcmap)
}

func (mcmap *Mcmap) Save() (rowsAffected int64) {
	result := database.DB.Save(&mcmap)
	return result.RowsAffected
}

func (mcmap *Mcmap) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&mcmap)
	return result.RowsAffected
}
