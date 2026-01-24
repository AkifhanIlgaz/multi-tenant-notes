package migrations

import (
	"fmt"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"gorm.io/gorm"
)

var tenants = []models.Tenant{
	{Name: "Beyaz Futbol", Slug: "beyaz_futbol"},
	{Name: "Hell Kitchen", Slug: "hell_kitchen"},
	{Name: "Mentalist", Slug: "mentalist"},
}

func RunMigrations(db *gorm.DB) error {
	for _, tenant := range tenants {
		if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", tenant.Slug)).Error; err != nil {
			return err
		}

		if err := db.Exec(fmt.Sprintf("SET search_path TO %s", tenant.Slug)).Error; err != nil {
			fmt.Printf("Failed to set schema for tenant %s: %v\n", tenant.Slug, err)
			return nil
		}

		err := db.AutoMigrate(
			&models.User{},
			&models.Announcement{},
		)
		if err != nil {
			return err
		}
	}

	return nil
}

func DropTables(db *gorm.DB) error {
	for _, tenant := range tenants {
		if err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", tenant.Slug)).Error; err != nil {
			return err
		}

		if err := db.Exec(fmt.Sprintf("SET search_path TO %s", tenant.Slug)).Error; err != nil {
			fmt.Printf("Failed to set schema for tenant %s: %v\n", tenant.Slug, err)
			return nil
		}

		err := db.Migrator().DropTable(
			&models.User{},
			&models.Announcement{},
		)
		if err != nil {
			return err
		}
	}

	return nil
}
