package migrations

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.Tenant{},
		&models.User{},
		&models.Announcement{},
	)
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.Tenant{},
		&models.Announcement{},
		&models.User{},
	)
}
