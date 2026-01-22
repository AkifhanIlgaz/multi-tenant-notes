package models

import (
	"github.com/AkifhanIlgaz/shared-schema-and-db/internal/core/entity"
)

type UserModel struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"not null;uniqueIndex:idx_tenant_email"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`

	TenantId int          `gorm:"not null;uniqueIndex:idx_tenant_email;index"`
	Tenant   *TenantModel `gorm:"foreignKey:TenantId;constraint:OnDelete:CASCADE"`

	// Relations
	Announcements []AnnouncementModel `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (UserModel) TableName() string {
	return "users"
}

func NewUser(user *entity.User) *UserModel {
	if user == nil {
		return nil
	}

	return &UserModel{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		TenantId: user.TenantId,
	}
}
