package main

import (
	"log"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db/config"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db/migrations"
)

func main() {
	cfg := config.NewDatabaseConfig()
	if cfg == nil {
		log.Fatalf("Failed to create database configuration")
	}

	db, err := db.ConnectPostgres(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	if err := migrations.RunMigrations(db); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	if err := migrations.ResetAndSeed(db); err != nil {
		log.Fatalf("Failed to run seeds: %v", err)
	}
}
