package api

import (
	"github.com/gofiber/fiber/v3"
)

type Router struct {
	app                 *fiber.App
	authHandler         *AuthHandler
	announcementHandler *AnnouncementHandler
	authMiddleware      *AuthMiddleware
}

func NewRouter(app *fiber.App, authHandler *AuthHandler, announcementHandler *AnnouncementHandler, authMiddleware *AuthMiddleware) *Router {
	return &Router{
		app:                 app,
		authHandler:         authHandler,
		announcementHandler: announcementHandler,
		authMiddleware:      authMiddleware,
	}
}

func (r *Router) SetupRoutes() {
	r.app.Get("/health", func(c fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	api := r.app.Group("/api")

	// Public routes
	auth := api.Group("/auth")
	auth.Post("/login", r.authHandler.Login)

	// Protected routes
	protected := api.Group("/", r.authMiddleware.JWTMiddleware())

	notes := protected.Group("/notes")

	notes.Get("", r.announcementHandler.GetAnnouncementsOfTenant)
	notes.Post("", r.announcementHandler.CreateAnnouncement)
	notes.Delete("/:id", r.announcementHandler.DeleteAnnouncement)
}
