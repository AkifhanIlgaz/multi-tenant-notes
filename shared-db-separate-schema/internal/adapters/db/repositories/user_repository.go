package repositories

import (
	"context"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/ports"
	"gorm.io/gorm"
)

type userRepository struct {
	schemaWrapper *SchemaWrapper
	db            *gorm.DB
}

func NewUserRepository(db *gorm.DB, schemaWrapper *SchemaWrapper) ports.UserRepository {
	return &userRepository{schemaWrapper: schemaWrapper, db: db}
}

func (r *userRepository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (models.User, error) {
	user := models.User{}
	fn := func(db *gorm.DB) error {
		return db.Where("email = ? AND password = ? ", email, password).First(&user).Error
	}

	if err := r.schemaWrapper.ExecuteWithSchema(ctx, fn); err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	users := []models.User{}
	fn := func(db *gorm.DB) error {
		return db.Find(&users).Error
	}

	if err := r.schemaWrapper.ExecuteWithSchema(ctx, fn); err != nil {
		return nil, err
	}

	return users, nil
}
