package models

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"not null;uniqueIndex:idx_tenant_email"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`

	TenantId int     `gorm:"not null;uniqueIndex:idx_tenant_email;index"`
	Tenant   *Tenant `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE"`

	// Relations
	Announcements []Announcement `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}
