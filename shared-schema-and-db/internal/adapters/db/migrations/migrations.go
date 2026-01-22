package migrations

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/models"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {

	return db.AutoMigrate(
		&models.TenantModel{},
		&models.UserModel{},
		&models.AnnouncementModel{},
	)
}

func DropTables(db *gorm.DB) error {
	return db.Migrator().DropTable(
		&models.AnnouncementModel{},
		&models.UserModel{},
		&models.TenantModel{},
	)
}
