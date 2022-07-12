package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type ServiceToken struct {
		models.BaseModel

		Service string `gorm:"type:varchar(255);unique_index"`
		Token   string `gorm:"type:varchar(255);unique_index"`
		Scopes  string `gorm:"type:varchar(255)"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&ServiceToken{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&ServiceToken{})
	}

	migrate.Add("2022_07_11_170051_create_service_tokens_table", up, down)
}
