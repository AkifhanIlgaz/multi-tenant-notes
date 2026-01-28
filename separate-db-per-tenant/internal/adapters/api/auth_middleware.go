package api

import (
	"context"
	"strings"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/service"
	"github.com/gofiber/fiber/v3"
)

type AuthMiddleware struct {
	authService *service.AuthService
}

func NewAuthMiddleware(authService *service.AuthService) *AuthMiddleware {
	return &AuthMiddleware{
		authService: authService,
	}
}

func (m *AuthMiddleware) JWTMiddleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "missing authorization header",
			})
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "invalid authorization header format",
			})
		}

		tokenString := parts[1]

		claims, err := m.authService.ValidateToken(tokenString)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		c.Locals("user_id", claims.UserID)

		db := claims.TenantSlug
		ctx := context.WithValue(c.Context(), "db", db)
		c.SetContext(ctx)

		return c.Next()
	}
}

// Helper functions
func GetUserID(c fiber.Ctx) int {
	return c.Locals("user_id").(int)
}
