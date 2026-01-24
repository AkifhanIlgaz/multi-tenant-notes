package ports

import (
	"context"

	"github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"
)

type UserRepository interface {
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (models.User, error)
	GetUsers(ctx context.Context) ([]models.User, error)
}
