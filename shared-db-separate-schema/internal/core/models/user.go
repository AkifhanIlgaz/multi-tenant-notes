package models

type User struct {
	Id       int    `gorm:"primaryKey;autoIncrement"`
	Email    string `gorm:"not null;uniqueIndex:idx_tenant_email"`
	Name     string `gorm:"not null"`
	Password string `gorm:"not null"`

	// Relations
	Announcements []Announcement `gorm:"foreignKey:UserId;constraint:OnDelete:CASCADE"`
}

func (User) TableName() string {
	return "users"
}
