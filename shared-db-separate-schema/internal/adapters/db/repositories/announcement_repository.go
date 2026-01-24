package repositories

import (
	"context"
	"errors"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/ports"
	"gorm.io/gorm"
)

type announcementRepository struct {
	db            *gorm.DB
	schemaWrapper *SchemaWrapper
}

func NewAnnouncementRepository(db *gorm.DB, schemaWrapper *SchemaWrapper) ports.AnnouncementRepository {
	return &announcementRepository{db: db, schemaWrapper: schemaWrapper}
}

func (r *announcementRepository) CreateAnnouncement(ctx context.Context, announcement models.Announcement) error {
	fn := func(db *gorm.DB) error {
		return db.Create(&announcement).Error
	}
	return r.schemaWrapper.ExecuteWithSchema(ctx, fn)
}

func (r *announcementRepository) GetAnnouncements(ctx context.Context) ([]models.Announcement, error) {
	var announcements []models.Announcement

	fn := func(db *gorm.DB) error {
		return db.Find(&announcements).Error
	}

	err := r.schemaWrapper.ExecuteWithSchema(ctx, fn)
	if err != nil {
		return nil, err
	}

	return announcements, nil
}

func (r *announcementRepository) DeleteAnnouncement(ctx context.Context, announcementId int, userId int) error {
	fn := func(db *gorm.DB) error {
		res := db.Where("id = ? AND user_id = ?", announcementId, userId).Delete(&models.Announcement{})

		if res.RowsAffected == 0 {
			return errors.New("announcement not found or unauthorized to delete")
		}

		return res.Error
	}

	err := r.schemaWrapper.ExecuteWithSchema(ctx, fn)
	if err != nil {
		return err
	}

	return nil
}
