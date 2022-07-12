//Package service_token 模型
package service_token

import (
	"gohub/app/models"
	"gohub/pkg/database"
)

type ServiceToken struct {
	models.BaseModel

	Service string `json:"service"`
	Token   string `json:"token"`
	Scopes  string `json:"scopes"`

	models.CommonTimestampsField
}

func (serviceToken *ServiceToken) Create() {
	database.DB.Create(&serviceToken)
}

func (serviceToken *ServiceToken) Save() (rowsAffected int64) {
	result := database.DB.Save(&serviceToken)
	return result.RowsAffected
}

func (serviceToken *ServiceToken) Delete() (rowsAffected int64) {
	result := database.DB.Delete(&serviceToken)
	return result.RowsAffected
}
