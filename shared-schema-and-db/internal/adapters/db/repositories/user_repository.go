package repositories

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmailAndPassword(email string, password string) (models.User, error) {
	user := models.User{}
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUsersByTenantId(tenantId int) ([]models.User, error) {
	users := []models.User{}
	if err := r.db.Where("tenant_id = ?", tenantId).Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
