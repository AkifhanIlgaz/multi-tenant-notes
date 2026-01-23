package service

import (
	"fmt"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/ports"
)

type AnnouncementService struct {
	repo ports.AnnouncementRepository
}

func NewAnnouncementService(repo ports.AnnouncementRepository) *AnnouncementService {
	return &AnnouncementService{
		repo: repo,
	}
}

func (s *AnnouncementService) GetAnnouncementsByTenantID(tenantId int) ([]models.Announcement, error) {
	announcements, err := s.repo.GetAnnouncementsByTenantId(tenantId)
	if err != nil {
		return nil, fmt.Errorf("failed to get announcements by tenant id: %w", err)
	}

	return announcements, nil
}

func (s *AnnouncementService) CreateAnnouncement(announcement models.Announcement) error {
	err := s.repo.CreateAnnouncement(announcement)
	if err != nil {
		return fmt.Errorf("failed to create announcement: %w", err)
	}

	return nil
}

func (s *AnnouncementService) DeleteAnnouncement(announcementId, userId, tenantId int) error {
	err := s.repo.DeleteAnnouncement(announcementId, userId, tenantId)
	if err != nil {
		return fmt.Errorf("failed to delete announcement: %w", err)
	}

	return nil
}
