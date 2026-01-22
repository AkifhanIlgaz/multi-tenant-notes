package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"

type UserRepository interface {
	GetUserByEmailAndPassword(email string, password string) (models.User, error)
	GetUsersByTenantId(tenantId int) ([]models.User, error)
}
