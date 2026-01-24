package ports

import "github.com/AkifhanIlgaz/shared-db-separate-schema/internal/core/models"

type TenantRepository interface {
	GetTenantBySlug(slug string) (models.Tenant, error)
	GetTenantById(id int) (models.Tenant, error)
}
