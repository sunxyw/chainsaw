package migrations

import (
	"database/sql"
	"gohub/app/models"
	"gohub/pkg/migrate"
	"time"

	"gorm.io/gorm"
)

func init() {

	type Notification struct {
		models.BaseModel

		Type   string    `gorm:"type:varchar(255)"`
		Data   string    `gorm:"type:varchar(255)"`
		ReadAt time.Time `gorm:"type:timestamp"`

		models.CommonTimestampsField
	}

	up := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.AutoMigrate(&Notification{})
	}

	down := func(migrator gorm.Migrator, DB *sql.DB) {
		migrator.DropTable(&Notification{})
	}

	migrate.Add("2022_07_14_164946_create_notifications_table", up, down)
}
