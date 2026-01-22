package models

import "time"

type Announcement struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`

	UserId int   `gorm:"not null;index"`
	User   *User `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	TenantId int     `gorm:"not null;index"`
	Tenant   *Tenant `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE"`
}

func (Announcement) TableName() string {
	return "announcements"
}
