package handler

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/api/dto"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/service"
	"github.com/gofiber/fiber/v3"
)

type AuthHandler struct {
	authService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{
		authService: authService,
	}
}

func (h *AuthHandler) Login(c fiber.Ctx) error {
	var req dto.LoginRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	user, err := h.authService.Login(req.Email, req.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	token, err := h.authService.GenerateToken(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	response := dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID:       user.Id,
			Email:    user.Email,
			Name:     user.Name,
			TenantID: user.TenantId,
		},
	}

	return c.Status(fiber.StatusOK).JSON(response)
}
