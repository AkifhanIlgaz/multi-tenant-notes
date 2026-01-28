package migrations

import (
	"context"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/db/repositories"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/models"
)

func RunMigrations(dbMux *repositories.DBMultiplexer) error {

	for _, tenant := range repositories.Tenants {
		db, err := dbMux.GetClient(context.WithValue(context.Background(), "db", tenant.Slug))
		if err != nil {
			return err
		}

		if err = db.AutoMigrate(
			&models.User{},
			&models.Announcement{},
		); err != nil {
			return err
		}
	}

	return nil
}

func DropTables(dbMux *repositories.DBMultiplexer) error {
	for _, tenant := range repositories.Tenants {
		db, err := dbMux.GetClient(context.WithValue(context.Background(), "db", tenant.Slug))
		if err != nil {
			return err
		}

		if err = db.Migrator().DropTable(
			&models.User{},
			&models.Announcement{},
		); err != nil {
			return err
		}
	}

	return nil
}
