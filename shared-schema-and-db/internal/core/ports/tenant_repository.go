package ports

import "github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"

type TenantRepository interface {
	GetTenantBySlug(slug string) (models.Tenant, error)
	GetTenantById(id int) (models.Tenant, error)
}
