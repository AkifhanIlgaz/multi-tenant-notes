package repositories

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/mappers"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/adapters/db/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByEmailAndPassword(email string, password string) (entity.User, error) {
	user := models.UserModel{}
	if err := r.db.Where("email = ? AND password = ?", email, password).First(&user).Error; err != nil {
		return entity.User{}, err
	}

	return mappers.ToUserEntity(user), nil
}

func (r *userRepository) GetUsersByTenantId(tenantId int) ([]entity.User, error) {
	users := []models.UserModel{}
	if err := r.db.Where("tenant_id = ?", tenantId).Find(&users).Error; err != nil {
		return nil, err
	}

	return mappers.ToUserEntities(users), nil
}
