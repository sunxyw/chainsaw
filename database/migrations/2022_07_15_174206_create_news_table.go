package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type New struct {
		models.BaseModel

		Title  string `gorm:"type:varchar(255);not null"`
		URL    string `gorm:"type:varchar(255);not null"`
		Source string `gorm:"type:varchar(255);not null"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&New{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&New{})
	}

	migrate.Add("2022_07_15_174206_create_news_table", up, down)
}
