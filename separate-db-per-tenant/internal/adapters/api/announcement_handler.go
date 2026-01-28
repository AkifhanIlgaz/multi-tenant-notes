package api

import (
	"strconv"
	"time"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/models"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/service"
	"github.com/gofiber/fiber/v3"
)

type AnnouncementHandler struct {
	service *service.AnnouncementService
}

func NewAnnouncementHandler(service *service.AnnouncementService) *AnnouncementHandler {
	return &AnnouncementHandler{
		service: service,
	}
}

func (h *AnnouncementHandler) GetAnnouncements(ctx fiber.Ctx) error {
	announcements, err := h.service.GetAnnouncements(ctx.Context())
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to retrieve announcements",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"announcements": announcements,
	})
}

func (h *AnnouncementHandler) CreateAnnouncement(ctx fiber.Ctx) error {

	var req struct {
		Title   string `json:"title"`
		Content string `json:"content"`
	}
	if err := ctx.Bind().Body(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "invalid request body",
		})
	}

	announcement := models.Announcement{
		Title:     req.Title,
		Content:   req.Content,
		CreatedAt: time.Now(),
		UserId:    GetUserID(ctx),
	}

	if err := h.service.CreateAnnouncement(ctx.Context(), announcement); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to create announcement",
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{
		"message": "announcement created successfully",
	})
}

func (h *AnnouncementHandler) DeleteAnnouncement(ctx fiber.Ctx) error {
	idStr := ctx.Params("id")

	if idStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "announcement ID is missing or invalid",
		})
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "announcement ID is missing or invalid",
		})
	}

	userID := GetUserID(ctx)

	if err := h.service.DeleteAnnouncement(ctx.Context(), id, userID); err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "announcement deleted successfully",
	})
}
