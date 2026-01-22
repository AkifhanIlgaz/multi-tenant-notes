package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"

type AnnouncementRepository interface {
	CreateAnnouncement(announcement models.Announcement) error
	GetAnnouncementsByTenantId(tenantId int) ([]models.Announcement, error)
	GetAnnouncementsByUserId(userId int) ([]models.Announcement, error)
	DeleteAnnouncement(id int) error
}
