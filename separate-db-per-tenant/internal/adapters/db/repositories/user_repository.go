package repositories

import (
	"context"

	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/models"
	"github.com/AkifhanIlgaz/separate-db-per-tenant/internal/core/ports"
)

type userRepository struct {
	*DBMultiplexer
}

func NewUserRepository(dbMux *DBMultiplexer) ports.UserRepository {
	return &userRepository{DBMultiplexer: dbMux}
}

func (r *userRepository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (models.User, error) {
	db, err := r.GetClient(ctx)
	if err != nil {
		return models.User{}, err
	}

	user := models.User{}
	if err := db.Where("email = ? AND password = ? ", email, password).First(&user).Error; err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]models.User, error) {
	db, err := r.GetClient(ctx)
	if err != nil {
		return []models.User{}, err
	}

	users := []models.User{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
