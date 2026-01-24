package service

import (
	"context"
	"fmt"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/ports"
)

type AnnouncementService struct {
	repo ports.AnnouncementRepository
}

func NewAnnouncementService(repo ports.AnnouncementRepository) *AnnouncementService {
	return &AnnouncementService{
		repo: repo,
	}
}

func (s *AnnouncementService) GetAnnouncements(ctx context.Context) ([]models.Announcement, error) {
	announcements, err := s.repo.GetAnnouncements(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements by tenant id: %w", err)
	}

	return announcements, nil
}

func (s *AnnouncementService) CreateAnnouncement(ctx context.Context, announcement models.Announcement) error {
	err := s.repo.CreateAnnouncement(ctx, announcement)
	if err != nil {
		return fmt.Errorf("failed to create announcement: %w", err)
	}

	return nil
}

func (s *AnnouncementService) DeleteAnnouncement(ctx context.Context, announcementId, userId int) error {
	err := s.repo.DeleteAnnouncement(ctx, announcementId, userId)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	return nil
}
