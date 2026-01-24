package migrations

import (
	"fmt"
	"strings"
	"time"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"gorm.io/gorm"
)

func SeedDatabase(db *gorm.DB) error {

	// Her tenant için user'lar ve announcement'lar oluştur
	for _, tenant := range tenants {
		if err := db.Exec(fmt.Sprintf("SET search_path TO %s", tenant.Slug)).Error; err != nil {
			fmt.Printf("Failed to set schema for tenant %s: %v\n", tenant.Slug, err)
			return nil
		}

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

func createUsersForTenant(db *gorm.DB, tenant models.Tenant) []models.User {
	var userNames []string

	// Her tenant için farklı user isimleri
	switch tenant.Name {
	case "Beyaz Futbol":
		userNames = []string{"sinan.engin", "ahmet.cakar", "ertem.sener"}
	case "Hell Kitchen":
		userNames = []string{"gordon.ramsay", "mehmet.yalcinkaya", "sofia.fehn"}
	case "Mentalist":
		userNames = []string{"patrick.jane", "kimball.cho", "teresa.lisbon"}
	default:
		userNames = []string{"john.doe", "jane.smith", "alex.wilson"}
	}

	users := make([]models.User, 0, len(userNames))

	for _, userName := range userNames {
		email := fmt.Sprintf("%s@%s.com", userName, tenant.Slug)

		user := models.User{
			Email:    email,
			Password: "password123", // Demo için plain text
			Name:     usernameToName(userName),
		}

		if err := db.Create(&user).Error; err != nil {
			fmt.Printf("Failed to create user %s for tenant %s: %v\n", userName, tenant.Slug, err)
			return nil
		}
		users = append(users, user)
	}

	return users
}

func createAnnouncementsForTenant(tenant models.Tenant, users []models.User) []models.Announcement {
	var announcements []models.Announcement

	switch tenant.Name {
	case "Beyaz Futbol":
		announcements = []models.Announcement{
			{
				Title:     "Yeni Sezon Başlıyor!",
				Content:   "Sevgili futbolseverler, yeni sezonda da sizlerle birlikte olmaktan mutluluk duyuyoruz. Bu sezon çok özel analizler ve sıcak tartışmalarla dolu bir sezon olacak. Heyecanla bekliyoruz!",
				CreatedAt: time.Now().Add(-72 * time.Hour),
				UserId:    users[0].Id,
			},
			{
				Title:     "Özel Canlı Yayın Duyurusu",
				Content:   "Bu hafta Cumartesi akşamı 20:00'de özel bir canlı yayınımız var! Haftanın maçlarını değerlendireceğiz ve sürpriz konuklarımız olacak. Kaçırmayın!",
				CreatedAt: time.Now().Add(-48 * time.Hour),
				UserId:    users[1].Id,
			},
			{
				Title:     "Ağaca Sormuşlar",
				Content:   "Kökü benden de ondan demiş. Abi dediğim adam yaptı bana bunu yaptı. Bana bi üç beş dakika müsaade edin çocuklar.",
				CreatedAt: time.Now().Add(-24 * time.Hour),
				UserId:    users[2].Id,
			},
		}

	case "Hell Kitchen":
		announcements = []models.Announcement{
			{
				Title:     "New Season Auditions Open!",
				Content:   "We're looking for the best chefs to compete in the next season of Hell's Kitchen! If you think you have what it takes to handle the heat, apply now. Gordon is waiting!",
				CreatedAt: time.Now().Add(-96 * time.Hour),
				UserId:    users[0].Id,
			},
			{
				Title:     "Michelin Yıldızlı Menü Hazırlığı",
				Content:   "Bu hafta mutfağımızda Michelin yıldızlı özel bir menü hazırlayacağız. Yarışmacılar en üst düzey tekniklerle test edilecek. Hazır olun!",
				CreatedAt: time.Now().Add(-60 * time.Hour),
				UserId:    users[1].Id,
			},
			{
				Title:     "Restaurant Service Challenge",
				Content:   "Next episode features our famous restaurant service challenge. Both teams will serve real customers under intense pressure. Remember: perfection is the only acceptable standard!",
				CreatedAt: time.Now().Add(-12 * time.Hour),
				UserId:    users[2].Id,
			},
		}

	case "Mentalist":
		announcements = []models.Announcement{
			{
				Title:     "New Case: The Red John Investigation",
				Content:   "CBI has reopened several cold cases connected to Red John. Patrick Jane will be leading the investigation with his unique methods. All team members report to briefing room at 0800 hours.",
				CreatedAt: time.Now().Add(-120 * time.Hour),
				UserId:    users[0].Id,
			},
			{
				Title:     "Mandatory Interrogation Training",
				Content:   "All agents must complete the advanced interrogation techniques course by end of month. Agent Cho will be conducting the sessions. Psychological profiling techniques will be covered.",
				CreatedAt: time.Now().Add(-36 * time.Hour),
				UserId:    users[1].Id,
			},
			{
				Title:     "Field Operation Protocol Update",
				Content:   "New protocols for field operations are now in effect. Agent Lisbon has updated all team procedures. Please review the documentation and sign off by Friday. Safety first.",
				CreatedAt: time.Now().Add(-6 * time.Hour),
				UserId:    users[2].Id,
			},
		}
	}

	return announcements
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
