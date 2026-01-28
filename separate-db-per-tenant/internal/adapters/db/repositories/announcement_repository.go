package repositories

import (
	"context"
	"errors"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/models"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/ports"
)

type announcementRepository struct {
	*DBMultiplexer
}

func NewAnnouncementRepository(dbMux *DBMultiplexer) ports.AnnouncementRepository {
	return &announcementRepository{DBMultiplexer: dbMux}
}

func (r *announcementRepository) CreateAnnouncement(ctx context.Context, announcement models.Announcement) error {
	db, err := r.GetClient(ctx)
	if err != nil {
		return err
	}

	if err := db.Create(&announcement).Error; err != nil {
		return err
	}
	return nil
}

func (r *announcementRepository) GetAnnouncements(ctx context.Context) ([]models.Announcement, error) {
	db, err := r.GetClient(ctx)
	if err != nil {
		return nil, err
	}

	var announcements []models.Announcement

	if err := db.Find(&announcements).Error; err != nil {
		return nil, err
	}

	return announcements, nil
}

func (r *announcementRepository) DeleteAnnouncement(ctx context.Context, announcementId int, userId int) error {
	db, err := r.GetClient(ctx)
	if err != nil {
		return err
	}

	res := db.Where("id = ? AND user_id = ?", announcementId, userId).Delete(&models.Announcement{})

	if res.RowsAffected == 0 {
		return errors.New("announcement not found or unauthorized to delete")
	}

	return res.Error
}
