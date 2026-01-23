package repositories

import (
	"errors"
	"fmt"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/ports"
	"gorm.io/gorm"
)

type announcementRepository struct {
	db *gorm.DB
}

func NewAnnouncementRepository(db *gorm.DB) ports.AnnouncementRepository {
	return &announcementRepository{db: db}
}

func (r *announcementRepository) CreateAnnouncement(announcement models.Announcement) error {
	return r.db.Create(&announcement).Error
}

func (r *announcementRepository) GetAnnouncementsByTenantId(tenantId int) ([]models.Announcement, error) {
	var announcements []models.Announcement
	if err := r.db.Where("tenant_id = ?", tenantId).Find(&announcements).Error; err != nil {
		return nil, err
	}
	return announcements, nil
}

func (r *announcementRepository) DeleteAnnouncement(announcementId int, userId int, tenantId int) error {
	res := r.db.Where("id = ? AND tenant_id = ? AND user_id = ?", announcementId, tenantId, userId).Delete(&models.Announcement{})

	if res.Error != nil {
		return fmt.Errorf("failed to delete announcement: %w", res.Error)
	}

	if res.RowsAffected == 0 {
		return errors.New("announcement not found or unauthorized to delete")
	}

	return nil
}
