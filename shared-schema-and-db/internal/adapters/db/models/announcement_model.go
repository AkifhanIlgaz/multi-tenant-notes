package models

import "time"

type AnnouncementModel struct {
	Id        int       `gorm:"primaryKey;autoIncrement"`
	Title     string    `gorm:"not null"`
	Content   string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"not null"`

	UserId int        `gorm:"not null;index"`
	User   *UserModel `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`

	TenantId int          `gorm:"not null;index"`
	Tenant   *TenantModel `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE"`
}

func (AnnouncementModel) TableName() string {
	return "announcements"
}
