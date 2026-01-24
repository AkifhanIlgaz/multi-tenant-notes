package main

import (
	"log"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/api"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db/config"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db/migrations"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/adapters/db/repositories"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/service"
	"github.com/gofiber/fiber/v3"
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

	schemaWrapper := repositories.NewSchemaWrapper(db)
	userRepo := repositories.NewUserRepository(db, schemaWrapper)
	announcementRepo := repositories.NewAnnouncementRepository(db, schemaWrapper)

	authService := service.NewAuthService(userRepo)
	announcementService := service.NewAnnouncementService(announcementRepo)

	authHandler := api.NewAuthHandler(authService)
	announcementHandler := api.NewAnnouncementHandler(announcementService)

	authMiddleware := api.NewAuthMiddleware(authService)

	app := fiber.New()

	router := api.NewRouter(app, authHandler, announcementHandler, authMiddleware)
	router.SetupRoutes()

	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
