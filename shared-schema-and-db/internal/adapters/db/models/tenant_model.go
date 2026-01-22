package models

type TenantModel struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;uniqueIndex:idx_tenant_name"`
}

func (TenantModel) TableName() string {
	return "tenants"
}
