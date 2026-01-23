package repositories

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/models"
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/ports"
	"gorm.io/gorm"
)

type tenantRepository struct {
	db *gorm.DB
}

func NewTenantRepository(db *gorm.DB) ports.TenantRepository {
	return &tenantRepository{db: db}
}

func (r *tenantRepository) GetTenantBySlug(slug string) (models.Tenant, error) {
	var tenant models.Tenant
	if err := r.db.Where("slug = ?", slug).First(&tenant).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}

func (r *tenantRepository) GetTenantById(id int) (models.Tenant, error) {
	var tenant models.Tenant
	if err := r.db.Where("id = ?", id).First(&tenant).Error; err != nil {
		return models.Tenant{}, err
	}
	return tenant, nil
}
