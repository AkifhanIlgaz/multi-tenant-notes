package ports

import (
	"context"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
)

type AnnouncementRepository interface {
	CreateAnnouncement(ctx context.Context, announcement models.Announcement) error
	GetAnnouncements(ctx context.Context) ([]models.Announcement, error)
	DeleteAnnouncement(ctx context.Context, announcementId int, userId int) error
}
