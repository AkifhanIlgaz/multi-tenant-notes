package main

import (
	"log"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/api"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/db/migrations"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/adapters/db/repositories"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/service"
	"github.com/gofiber/fiber/v3"
)

func main() {
	dbMux, err := repositories.NewDBMultiplexer()
	if err != nil {
		log.Fatalf("Failed to create database multiplexer: %v", err)
	}

	if err := migrations.RunMigrations(dbMux); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}
	if err := migrations.ResetAndSeed(dbMux); err != nil {
		log.Fatalf("Failed to run seeds: %v", err)
	}

	userRepo := repositories.NewUserRepository(dbMux)
	announcementRepo := repositories.NewAnnouncementRepository(dbMux)

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
