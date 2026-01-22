package api

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/handler"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/middleware"
	"github.com/gofiber/fiber/v3"
)

func SetupRoutes(app *fiber.App, authHandler *handler.AuthHandler, authMiddleware *middleware.AuthMiddleware) {
	api := app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/login", authHandler.Login)

	// Protected routes
	protected := api.Group("/", authMiddleware.JWTMiddleware())
	protected.Get("/me", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"user_id":   middleware.GetUserID(c),
			"tenant_id": middleware.GetTenantID(c),
		})
	})

	// Health check
	app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})
}
