package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"

type UserRepository interface {
	GetUserByEmailAndPassword(email string, password string) (entity.User, error)
	GetUsersByTenantId(tenantId int) ([]entity.User, error)
}
