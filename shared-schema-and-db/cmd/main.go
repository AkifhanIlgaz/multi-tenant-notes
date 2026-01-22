package main

import (
	"log"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/handler"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/middleware"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/config"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/migrations"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/repositories"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/service"
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

	userRepo := repositories.NewUserRepository(db)

	authService := service.NewAuthService(userRepo)

	authHandler := handler.NewAuthHandler(authService)

	authMiddleware := middleware.NewAuthMiddleware(authService)

	app := fiber.New()

	api.SetupRoutes(app, authHandler, authMiddleware)

	log.Println("Server starting on :3000")
	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
