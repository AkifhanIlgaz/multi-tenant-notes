package models

type Tenant struct {
	Id   int    `gorm:"primaryKey;autoIncrement"`
	Name string `gorm:"not null;uniqueIndex:idx_tenant_name"`
}

func (Tenant) TableName() string {
	return "tenants"
}
