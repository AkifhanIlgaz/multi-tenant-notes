package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"

type TenantRepository interface {
	CreateTenant(tenant *entity.Tenant) error
}
