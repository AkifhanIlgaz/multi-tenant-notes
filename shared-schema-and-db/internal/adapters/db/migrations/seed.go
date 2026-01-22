package migrations

import (
	"fmt"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {
	// Tenant'ları oluştur
	tenants := []models.TenantModel{
		{Name: "Acme Corp"},
		{Name: "TechStart Inc"},
		{Name: "Global Solutions"},
	}

	for i := range tenants {
		if err := db.Create(&tenants[i]).Error; err != nil {
			return fmt.Errorf("failed to create tenant: %w", err)
		}
	}

	// Her tenant için user'lar ve announcement'lar oluştur
	for _, tenant := range tenants {
		// Her tenant için 3 user oluştur
		users := createUsersForTenant(db, tenant)
		if users == nil {
			return fmt.Errorf("failed to create users for tenant: %s", tenant.Name)
		}

		// Her tenant için genele hitap eden announcement'lar oluştur
		announcements := createAnnouncementsForTenant(tenant, users)
		for i := range announcements {
			if err := db.Create(&announcements[i]).Error; err != nil {
				return fmt.Errorf("failed to create announcement: %w", err)
			}
		}
	}

	return nil
}

func createUsersForTenant(db *gorm.DB, tenant models.TenantModel) []models.UserModel {
	userNames := []string{"john.doe", "jane.smith", "alex.wilson"}
	users := make([]models.UserModel, 0, len(userNames))

	for _, userName := range userNames {
		email := fmt.Sprintf("%s@%s.com", userName, slugify(tenant.Name))

		user := models.UserModel{
			Email:    email,
			Password: "password123", // Demo için plain text
			Name:     usernameToName(userName),
			TenantId: tenant.Id,
		}

		if err := db.Create(&user).Error; err != nil {
			return nil
		}
		users = append(users, user)
	}

	return users
}

func createAnnouncementsForTenant(tenant models.TenantModel, users []models.UserModel) []models.AnnouncementModel {
	var announcements []models.AnnouncementModel

	switch tenant.Name {
	case "Acme Corp":
		announcements = []models.AnnouncementModel{
			{
				Title:     "Welcome to Acme Corp!",
				Content:   "We're thrilled to welcome all team members to Acme Corp. Our mission is to deliver innovative solutions that transform businesses. Let's make this quarter our best one yet!",
				CreatedAt: time.Now().Add(-72 * time.Hour),
				UserId:    users[0].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "Q4 Company All-Hands Meeting",
				Content:   "Join us this Friday at 2 PM for our quarterly all-hands meeting. We'll be discussing our achievements, upcoming projects, and celebrating our team's success. Pizza and refreshments will be provided!",
				CreatedAt: time.Now().Add(-48 * time.Hour),
				UserId:    users[1].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "New Office Security Protocols",
				Content:   "Effective immediately, all employees must use their access badges to enter the building. Please ensure you have your badge with you at all times. Contact HR if you need a replacement.",
				CreatedAt: time.Now().Add(-24 * time.Hour),
				UserId:    users[2].Id,
				TenantId:  tenant.Id,
			},
		}

	case "TechStart Inc":
		announcements = []models.AnnouncementModel{
			{
				Title:     "TechStart Inc Holiday Schedule",
				Content:   "Our offices will be closed from December 24th through January 2nd for the holiday season. Emergency support will be available via email. Wishing everyone a wonderful holiday!",
				CreatedAt: time.Now().Add(-96 * time.Hour),
				UserId:    users[0].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "Launch of New Product Development Initiative",
				Content:   "We're excited to announce our new product development initiative! This strategic move will position us as industry leaders. All departments are encouraged to submit innovative ideas through the company portal.",
				CreatedAt: time.Now().Add(-60 * time.Hour),
				UserId:    users[1].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "Team Building Event Next Month",
				Content:   "Mark your calendars! We're organizing a team building retreat next month at Mountain View Resort. Activities include hiking, workshops, and networking sessions. RSVP by end of week.",
				CreatedAt: time.Now().Add(-12 * time.Hour),
				UserId:    users[2].Id,
				TenantId:  tenant.Id,
			},
		}

	case "Global Solutions":
		announcements = []models.AnnouncementModel{
			{
				Title:     "Global Solutions Expands to Three New Markets",
				Content:   "We're proud to announce our expansion into Asia, South America, and Africa! This milestone represents years of hard work and dedication. Thank you to everyone who made this possible.",
				CreatedAt: time.Now().Add(-120 * time.Hour),
				UserId:    users[0].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "Mandatory Cybersecurity Training",
				Content:   "All employees must complete the cybersecurity awareness training by end of month. This training covers phishing detection, password security, and data protection. Access the course through our learning portal.",
				CreatedAt: time.Now().Add(-36 * time.Hour),
				UserId:    users[1].Id,
				TenantId:  tenant.Id,
			},
			{
				Title:     "Employee Wellness Program Launch",
				Content:   "We're launching a comprehensive wellness program including gym memberships, mental health resources, and flexible work arrangements. Visit the HR portal to learn more and sign up for benefits.",
				CreatedAt: time.Now().Add(-6 * time.Hour),
				UserId:    users[2].Id,
				TenantId:  tenant.Id,
			},
		}
	}

	return announcements
}

func slugify(s string) string {
	result := ""
	for _, char := range s {
		if char == ' ' {
			continue
		}
		if char >= 'A' && char <= 'Z' {
			result += string(char + 32)
		} else {
			result += string(char)
		}
	}
	return result
}

func usernameToName(username string) string {
	parts := strings.Split(username, ".")
	titleCaser := cases.Title(language.English)

	firstName := titleCaser.String(parts[0])
	lastName := titleCaser.String(parts[1])
	return fmt.Sprintf("%s %s", firstName, lastName)
}

func ResetAndSeed(db *gorm.DB) error {
	if err := DropTables(db); err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	if err := RunMigrations(db); err != nil {
		return fmt.Errorf("failed to run migrations: %w", err)
	}

	if err := SeedDatabase(db); err != nil {
		return fmt.Errorf("failed to seed database: %w", err)
	}

	return nil
}
