package models

import "time"

type Announcement struct {
	Id        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	Title     string    `gorm:"not null" json:"title"`
	Content   string    `gorm:"not null" json:"content"`
	CreatedAt time.Time `gorm:"not null" json:"created_at"`

	UserId int   `gorm:"not null;index" json:"user_id"`
	User   *User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE" json:"user,omitempty"`

	TenantId int     `gorm:"not null;index" json:"tenant_id"`
	Tenant   *Tenant `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE" json:"tenant,omitempty"`
}

func (Announcement) TableName() string {
	return "announcements"
}
