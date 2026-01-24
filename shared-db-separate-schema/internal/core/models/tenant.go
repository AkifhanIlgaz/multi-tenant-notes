package models

type Tenant struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;uniqueIndex:idx_tenant_name"`
	Slug string `gorm:"not null;uniqueIndex:idx_tenant_slug"`
}

func (Tenant) TableName() string {
	return "tenants"
}
