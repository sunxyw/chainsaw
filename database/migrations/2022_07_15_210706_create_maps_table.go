package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"

	"gorm.io/gorm"
)

func init() {

	type Map struct {
		models.BaseModel

		Name    string `gorm:"type:varchar(255);unique_index"`
		Type    string `gorm:"type:varchar(255)"`
		Address string `gorm:"type:varchar(255)"`
		Size    int64  `gorm:"type:bigint"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Map{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Map{})
	}

	migrate.Add("2022_07_15_210706_create_maps_table", up, down)
}
