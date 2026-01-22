package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"

type AnnouncementRepository interface {
	CreateAnnouncement(announcement *entity.Announcement) error
	GetAnnouncementsByTenantId(tenantId int) ([]*entity.Announcement, error)
	GetAnnouncementsByUserId(userId int) ([]*entity.Announcement, error)
	DeleteAnnouncement(id int) error
}
